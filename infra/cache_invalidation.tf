resource "aws_cloudwatch_event_rule" "posts_updated_rule" {
  name           = "posts-updated"
  event_bus_name = data.aws_cloudwatch_event_bus.default.name
  event_pattern = jsonencode({
    source = [
      "Pipe ${aws_pipes_pipe.posts_table_update.name}"
    ]
  })
}

resource "aws_cloudwatch_event_target" "cache_invalidator_lambda" {
  rule      = aws_cloudwatch_event_rule.posts_updated_rule.name
  arn       = aws_lambda_function.cache_invalidator.arn
}

# Lambda function to invalidate the cache

# Prepare lambda code

# data "archive_file" "cache_invalidator" {
#   type        = "zip"
#   source_file = "../src/cache-invalidator/bootstrap"
#   output_path = "../src/cache-invalidator/bootstrap.zip"
# }

# # Upload lambda code to S3

# resource "aws_s3_object" "cache_invalidator" {
#   bucket = aws_s3_bucket.lambda_bucket.id
#   key    = filemd5(data.archive_file.cache_invalidator.output_path)
#   source = data.archive_file.cache_invalidator.output_path
#   etag   = filemd5(data.archive_file.cache_invalidator.output_path)
# }

# resource "aws_lambda_function" "cache_invalidator" {
#   function_name = "invalidate-cache"

#   s3_bucket        = aws_s3_bucket.lambda_bucket.id
#   s3_key           = aws_s3_object.cache_invalidator.key
#   source_code_hash = data.archive_file.cache_invalidator.output_base64sha256

#   handler       = "bootstrap"
#   runtime       = "provided.al2"
#   architectures = ["arm64"]

#   role = aws_iam_role.cache_invalidator.arn
# }


resource "aws_lambda_function" "cache_invalidator" {
    function_name = "cache-invalidator"

    filename = "foo.zip"
    handler = "foo.handler"
    runtime = "python3.8"
    architectures = [ "arm64" ]

    role = aws_iam_role.cache_invalidator.arn
}

resource "aws_cloudwatch_log_group" "cache_invalidator" {
  name              = "/aws/lambda/cache-invalidator"
  retention_in_days = 14
}

resource "aws_iam_role" "cache_invalidator" {
  name = "cache_invalidator"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "invalidate_cloudfront_cache" {
  name = "invalidate_cloudfront_cache"
  role = aws_iam_role.cache_invalidator.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "cloudfront:CreateInvalidation"
        ]
        Effect   = "Allow"
        Resource = aws_cloudfront_distribution.blog_frontend.arn
      }
    ]
  })
}

resource "aws_lambda_permission" "allow_eventbridge" {
  statement_id  = "AllowExecutionFromEventBridge"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cache_invalidator.arn
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.posts_updated_rule.arn
}