terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}

data "aws_partition" "current" {}
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

locals {
  role_name = "${var.name}-lambda-execution-role"
}

resource "aws_lambda_function" "this" {
  function_name = var.name

  role = aws_iam_role.lambda.arn

  package_type  = "Image"
  architectures = ["arm64"]
  image_uri     = var.image_uri
  image_config {
    entry_point       = var.image_config_entry_point
    command           = var.image_config_command
    working_directory = var.image_config_working_directory
  }
}

#######
# IAM #
#######

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    condition {
      test     = "StringEquals"
      variable = "aws:SourceArn"
      values   = ["arn:${data.aws_partition.current}:lambda:${data.aws_region.current}:${data.aws_caller_identity.current}:function:${local.role_name}"]
    }
  }
}

resource "aws_iam_role" "lambda" {
  name               = local.role_name
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

###################
# CloudWatch Logs #
###################

resource "aws_cloudwatch_log_group" "lambda" {
  name              = "/aws/lambda/${var.name}"
  retention_in_days = var.cloudwatch_logs_retention_in_days
}

resource "aws_iam_role_policy" "logs" {
  name = "cloudwatch_logs_write"
  role = aws_iam_role.lambda.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogStream",
          "logs:PutLogEvents",
        ]
        Effect   = "Allow"
        Resource = aws_cloudwatch_log_group.lambda.arn
      },
    ]
  })
}

# Grant read write access to DynamoDB table
resource "aws_iam_role_policy" "dynamodb_read" {
  count = var.dynamodb_table_arn != null ? 1 : 0
  name  = "dynamodb_read"
  role  = aws_iam_role.lambda.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "dynamodb:BatchGetItem",
          "dynamodb:GetItem",
          "dynamodb:Query",
          "dynamodb:Scan",
          "dynamodb:BatchWriteItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:DeleteItem"
        ]
        Effect   = "Allow"
        Resource = var.dynamodb_table_arn
      },
    ]
  })
}
