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
  role_name = "${var.name}-lambda-exec"
}

resource "aws_lambda_function" "this" {
  function_name = var.name

  role = aws_iam_role.lambda.arn

  package_type  = var.image_uri != null ? "Image" : "Zip"
  architectures = ["arm64"]

  filename  = var.image_uri == null && var.filename != null ? var.filename : null
  source_code_hash = var.image_uri == null && var.filename != null ? filebase64sha256(var.filename) : null
  handler   = var.image_uri == null && var.handler != null ? var.handler : null
  runtime   = var.image_uri == null && var.runtime != null? var.runtime : null
  image_uri = var.image_uri != null ? var.image_uri : null

  dynamic "image_config" {
    for_each = var.image_config_entry_point != null || var.image_config_command != null || var.image_config_working_directory != null ? [1] : []
    content {
      entry_point       = var.image_config_entry_point
      command           = var.image_config_command
      working_directory = var.image_config_working_directory
    }
  }

  environment {
    variables = merge(var.environment_variables,{
      DYNAMODB_TABLE_NAME = var.dynamodb_table_name
    })
  }
}

#######
# IAM #
#######

resource "aws_iam_role" "lambda" {
  name = local.role_name
  assume_role_policy = var.iam_assume_role_policy != null ? var.iam_assume_role_policy : jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
        Condition = {
          StringEquals = {
            "aws:SourceAccount" : data.aws_caller_identity.current.account_id
            "aws:SourceArn" : "arn:${data.aws_partition.current.id}:lambda:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:function:${var.name}"
          }
        }
      }
    ]
  })
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
        Resource = "${aws_cloudwatch_log_group.lambda.arn}:*"
      },
    ]
  })
}

###############
# Permissions #
###############

# Grant read write access to DynamoDB table
resource "aws_iam_role_policy" "dynamodb_read_write" {
  count = var.has_dynamodb_table ? 1 : 0
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
        Resource = [
          var.dynamodb_table_arn,
          "${var.dynamodb_table_arn}/*"
        ]
      },
    ]
  })
}

# Add inline policies
resource "aws_iam_role_policy" "inline_policies" {
  for_each = var.iam_role_policies
  name     = each.key
  role     = aws_iam_role.lambda.id
  policy   = each.value
}
