terraform {
  required_version = "1.5.4"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}


resource "aws_sqs_queue" "this" {
  name                       = var.name
  visibility_timeout_seconds = 90
}

resource "aws_sqs_queue_policy" "policy" {
  for_each  = var.sqs_queue_policies
  queue_url = aws_sqs_queue.this.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = each.key
        Effect = "Allow"
        Principal = {
          Service = each.value.servicePrincipal
        }
        Action   = each.value.actions
        Resource = aws_sqs_queue.this.arn
        Condition = {
          ArnEquals = {
            "aws:SourceArn" : each.value.sourceArn
          }
        }
      }
    ]
  })
}

module "lambda" {
  source = "./../lambda"

  name = var.name

  filename = var.lambda_filename
  handler  = var.lambda_handler
  runtime  = var.lambda_runtime

  image_uri = var.lambda_image_uri

  iam_role_policies = merge(var.lambda_iam_role_policies, {
    "sqs" = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Action = [
            "sqs:ChangeMessageVisibility",
            "sqs:DeleteMessage",
            "sqs:GetQueueAttributes",
            "sqs:ReceiveMessage"
          ]
          Effect   = "Allow"
          Resource = aws_sqs_queue.this.arn
        },
      ]
    })
  })

  iam_assume_role_policy = var.lambda_iam_assume_role_policy != null ? var.lambda_iam_assume_role_policy : jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

  environment_variables = var.lambda_environment_variables
}

resource "aws_lambda_event_source_mapping" "object_updated" {
  event_source_arn                   = aws_sqs_queue.this.arn
  function_name                      = module.lambda.name
  batch_size                         = var.batch_size
  maximum_batching_window_in_seconds = var.max_batching_window_in_seconds
  scaling_config {
    maximum_concurrency = var.max_concurrency
  }
}
