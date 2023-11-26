resource "aws_dynamodb_table" "posts_table" {
  name                        = var.posts_table_name
  billing_mode                = "PAY_PER_REQUEST"
  hash_key                    = "id"
  range_key                   = "createdAt"
  deletion_protection_enabled = true

  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "createdAt"
    type = "N"
  }

  attribute {
    name = "slug"
    type = "S"
  }

  global_secondary_index {
    name            = "slug-index"
    hash_key        = "slug"
    projection_type = "ALL"
  }
}



# EventBridge pipe to publish update events to the event bus
resource "aws_pipes_pipe" "posts_table_update" {
  name    = local.forward_posts_table_update_events_pipe_name
  role_arn = aws_iam_role.posts_table_update_pipe_execution_role.arn
  source   = aws_dynamodb_table.posts_table.stream_arn
  target   = data.aws_cloudwatch_event_bus.default.arn

  source_parameters {
    dynamodb_stream_parameters {
      starting_position = "LATEST"
    }
  }
}

data "aws_cloudwatch_event_bus" "default" {
  name = "default"
}

data "aws_caller_identity" "main" {}

locals {
  forward_posts_table_update_events_pipe_name = "forward-posts-table-update-events"
}

resource "aws_iam_role" "posts_table_update_pipe_execution_role" {
  name = "posts-table-update-pipe-execution-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "pipes.amazonaws.com"
        },
        Condition = {
          StringEquals = {
            "aws:SourceAccount" : data.aws_caller_identity.main.account_id
            "aws:SourceArn" : "arn:aws:pipes:${var.aws_region}:${data.aws_caller_identity.main.account_id}:pipe/${local.forward_posts_table_update_events_pipe_name}"
          }
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "posts_table_update_pipe_execution_role_ddb_read_policy" {
  name = "DynamoDBStreamRead"
  role = aws_iam_role.posts_table_update_pipe_execution_role.id

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
        Resource = aws_dynamodb_table.posts_table.stream_arn
      }
    ]
  })
}

resource "aws_iam_role_policy" "posts_table_update_pipe_execution_role_eventbridge_write_policy" {
  name = "EventBridgeWrite"
  role = aws_iam_role.posts_table_update_pipe_execution_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "events:PutEvents"
        ]
        Effect   = "Allow"
        Resource = data.aws_cloudwatch_event_bus.default.arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "basic_execution_policy_cache_invalidator" {
  role       = aws_iam_role.cache_invalidator.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

