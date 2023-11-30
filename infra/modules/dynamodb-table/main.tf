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
  role_name = "${var.name}-forward-ddb-cdc-pipe-execution-role"
}

resource "aws_dynamodb_table" "this" {
  name         = var.name
  billing_mode = "PAY_PER_REQUEST"

  hash_key  = var.hash_key
  range_key = var.range_key

  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  point_in_time_recovery {
    enabled = true
  }

  dynamic "attribute" {
    for_each = var.attributes

    content {
      name = attribute.value.name
      type = attribute.value.type
    }
  }

  dynamic "global_secondary_index" {
    for_each = var.global_secondary_indexes

    content {
      name               = global_secondary_index.value.name
      hash_key           = global_secondary_index.value.hash_key
      projection_type    = global_secondary_index.value.projection_type
      range_key          = lookup(global_secondary_index.value, "range_key", null)
      read_capacity      = lookup(global_secondary_index.value, "read_capacity", null)
      write_capacity     = lookup(global_secondary_index.value, "write_capacity", null)
      non_key_attributes = lookup(global_secondary_index.value, "non_key_attributes", null)
    }
  }
}

# EventBridge bus where cdc events will be published
resource "aws_cloudwatch_event_bus" "cdc_bus" {
  count = var.expose_cdc_events && var.eventbridge_bus_name == null ? 1 : 0
  name  = var.name
}

data "aws_cloudwatch_event_bus" "cdc_bus" {
  count = var.expose_cdc_events && var.eventbridge_bus_name != null ? 1 : 0
  name  = var.eventbridge_bus_name
}

locals {
  eventbridge_bus_arn = var.expose_cdc_events && var.eventbridge_bus_name != null ? data.aws_cloudwatch_event_bus.cdc_bus[0].arn : aws_cloudwatch_event_bus.cdc_bus[0].arn
  eventbridge_bus_name = var.expose_cdc_events && var.eventbridge_bus_name != null ? var.eventbridge_bus_name : aws_cloudwatch_event_bus.cdc_bus[0].name
}
# EventBridge pipe to forward update events to the event bus
resource "aws_pipes_pipe" "forward_cdc" {
  count    = var.expose_cdc_events ? 1 : 0
  name     = var.name
  role_arn = aws_iam_role.pipe[0].arn
  source   = aws_dynamodb_table.this.stream_arn
  target   = local.eventbridge_bus_arn

  source_parameters {
    dynamodb_stream_parameters {
      starting_position = "LATEST"
    }
  }
}


# IAM role for the pipe

resource "aws_iam_role" "pipe" {
  count = var.expose_cdc_events ? 1 : 0
  name  = local.role_name

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "pipes.amazonaws.com"
        },
        Condition = {
          StringEquals = {
            "aws:SourceAccount" : data.aws_caller_identity.current.account_id
            "aws:SourceArn" : "arn:aws:pipes:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:pipe/${var.name}"
          }
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "pipe_execution_role_ddb_read" {
  count = var.expose_cdc_events ? 1 : 0
  name  = "DynamoDBStreamRead"
  role  = aws_iam_role.pipe[0].id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "dynamodb:DescribeStream",
          "dynamodb:GetRecords",
          "dynamodb:GetShardIterator",
          "dynamodb:ListStreams"
        ]
        Effect   = "Allow"
        Resource = aws_dynamodb_table.this.stream_arn
      },
    ]
  })
}

resource "aws_iam_role_policy" "pipe_execution_role_eventbridge_write" {
  count = var.expose_cdc_events ? 1 : 0
  name  = "EventBridgeWrite"
  role  = aws_iam_role.pipe[0].id
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
