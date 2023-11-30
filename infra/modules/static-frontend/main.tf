terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

locals {
  eventbridge_bus_arn       = var.eventbridge_bus_name != null ? data.aws_cloudwatch_event_bus.event_bus[0].arn : aws_cloudwatch_event_bus.event_bus[0].arn
  eventbridge_bus_name      = var.eventbridge_bus_name != null ? var.eventbridge_bus_name : aws_cloudwatch_event_bus.event_bus[0].name
  pipe_name                 = "${var.name}-event-batching"
  batching_lambda_image_uri = ""
}

module "s3_bucket" {
  source                        = "./../s3-bucket"
  name                          = var.name
  publish_events_on_eventbridge = true
}

resource "aws_sqs_queue" "event_batching" {
  name = "${var.name}-event-batching"
}

# Allow Eventbridge to publish events to the SQS queue
resource "aws_sqs_queue_policy" "allow_eventbridge" {
  queue_url = aws_sqs_queue.event_batching.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowEventBridge"
        Effect = "Allow"
        Principal = {
          Service = "events.amazonaws.com"
        }
        Action   = "sqs:SendMessage"
        Resource = aws_sqs_queue.event_batching.arn
        Condition = {
          ArnEquals = {
            "aws:SourceArn" : aws_cloudwatch_event_rule.object_updated.arn
          }
        }
      }
    ]
  })
}

resource "aws_cloudwatch_event_rule" "object_updated" {
  name           = "${var.name}-object-updated"
  event_bus_name = "default" # S3 bucket events are always published on the default event bus
  event_pattern = jsonencode({
    "source" : ["aws.s3"],
    "detail-type" : ["Object Created", "Object Deleted", "Object Restore Completed"],
    "detail" : {
      "bucket" : {
        "name" : [module.s3_bucket.name]
      }
    }
  })
}

resource "aws_cloudwatch_event_target" "object_updated" {
  rule = aws_cloudwatch_event_rule.object_updated.name
  arn  = aws_sqs_queue.event_batching.arn
}

# EventBridge bus where objects update events will be published
resource "aws_cloudwatch_event_bus" "event_bus" {
  count = var.eventbridge_bus_name == null ? 1 : 0
  name  = var.name
}

data "aws_cloudwatch_event_bus" "event_bus" {
  count = var.eventbridge_bus_name != null ? 1 : 0
  name  = var.eventbridge_bus_name
}

resource "aws_pipes_pipe" "event_batching" {
  name       = local.pipe_name
  role_arn   = aws_iam_role.pipe.arn
  source     = aws_sqs_queue.event_batching.arn
  target     = local.eventbridge_bus_arn
  enrichment = module.event_batching_lambda.lambda_function_arn

  source_parameters {
    # https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-sqs.html#:~:text=By%20default%2C%20EventBridge,size%20is%20reached.
    sqs_queue_parameters {
      batch_size                         = 10 # SQS supports up to 10 messages per batch
      maximum_batching_window_in_seconds = 10 # SQS supports up to 300 seconds (5 minutes) of batching
    }
  }
}

# Used external module to have inline code (didn't want to support this option on ./../lambda module)
module "event_batching_lambda" {
  source = "terraform-aws-modules/lambda/aws"

  function_name = "${var.name}-event-batching"
  handler       = "index.handler"
  runtime       = "python3.11"

  source_path = "${path.module}/assets/event-batching-lambda"

}

# Allow lambda to be triggered by the pipe
resource "aws_lambda_permission" "pipe" {
  statement_id  = "AllowExecutionFromPipe"
  action        = "lambda:InvokeFunction"
  function_name = module.event_batching_lambda.lambda_function_name
  principal     = "pipes.amazonaws.com"

  source_arn = aws_pipes_pipe.event_batching.arn
}

# IAM role for the pipe
resource "aws_iam_role" "pipe" {
  name = "${var.name}-pipe"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "pipes.amazonaws.com"
        }
        Condition = {
          StringEquals = {
            "aws:SourceAccount" : data.aws_caller_identity.current.account_id
            "aws:SourceArn" : "arn:aws:pipes:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:pipe/${local.pipe_name}"
          }
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "pipe_execution_role_sqs_read" {
  name = "SQSRead"
  role = aws_iam_role.pipe.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "sqs:ReceiveMessage",
          "sqs:DeleteMessage",
          "sqs:GetQueueAttributes"
        ],
        Effect   = "Allow",
        Resource = aws_sqs_queue.event_batching.arn
      }
    ]
  })
}

resource "aws_iam_role_policy" "pipe_execution_role_eventbridge_write" {
  name = "EventBridgeWrite"
  role = aws_iam_role.pipe.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "events:PutEvents"
        ]
        Effect   = "Allow"
        Resource = local.eventbridge_bus_arn
      },
    ]
  })
}

resource "aws_iam_role_policy" "pipe_execution_role_lambda_invoke" {
  name = "LambdaInvoke"
  role = aws_iam_role.pipe.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "lambda:InvokeFunction"
        ]
        Effect   = "Allow"
        Resource = module.event_batching_lambda.lambda_function_arn
      },
    ]
  })
}
