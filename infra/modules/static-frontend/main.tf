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
}

module "s3_bucket" {
  source                        = "./../s3-bucket"
  name                          = var.name
  publish_events_on_eventbridge = true
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
  arn  = module.deploy_event_publisher_lambda.queue_arn
}

# Lambda function that will be triggered by the SQS queue
module "deploy_event_publisher_lambda" {
  source = "./../sqs-lambda"

  name = "${var.name}-deploy-event-publisher"

  lambda_filename = data.archive_file.lambda_zip.output_path
  lambda_handler  = "index.handler"
  lambda_runtime  = "python3.11"

  lambda_iam_role_policies = {
    "eventbridge" = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Action = [
            "events:PutEvents",
          ]
          Effect   = "Allow"
          Resource = local.eventbridge_bus_arn
        },
      ]
    })
  }

  sqs_queue_policies = {
    "allowEventbridge" = {
        servicePrincipal = "events.amazonaws.com"
        actions          = ["sqs:SendMessage"]
        sourceArn        = aws_cloudwatch_event_rule.object_updated.arn
    }
  }

  lambda_environment_variables = {
    EVENTBRIDGE_BUS_NAME = local.eventbridge_bus_name,
    EVENT_SOURCE_NAME = var.name,
  }
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "${path.module}/assets/publish-deploy-event/index.py"
  output_path = "${path.module}/assets/publish-deploy-event.zip"
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
