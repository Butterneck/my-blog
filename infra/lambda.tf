# S3 bucket for lambda code

resource "random_pet" "lambda_bucket_name" {
  prefix = "butterneck-blog-functions"
  length = 4
}

resource "aws_s3_bucket" "lambda_bucket" {
  bucket = random_pet.lambda_bucket_name.id
}

# Prepare lambda code

data "archive_file" "blog_backend" {
  type        = "zip"
  source_file = "../src/bootstrap"
  output_path = "../src/bootstrap.zip"
}

# Upload lambda code to S3

resource "aws_s3_object" "blog_backend" {
  bucket = aws_s3_bucket.lambda_bucket.id
  key    = filemd5(data.archive_file.blog_backend.output_path)
  source = data.archive_file.blog_backend.output_path
  etag   = filemd5(data.archive_file.blog_backend.output_path)
}


# Lambda function

resource "aws_lambda_function" "blog_backend" {
  function_name = "blog-backend"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_object.blog_backend.key

  runtime       = "provided.al2"
  handler       = "bootstrap"
  architectures = ["arm64"]

  source_code_hash = data.archive_file.blog_backend.output_base64sha256

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_cloudwatch_log_group" "blog_backend" {
  name              = "/aws/lambda/blog-backend"
  retention_in_days = 14
}

resource "aws_iam_role" "lambda_exec" {
  name = "lambda_exec"
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

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_policy" "lambda_dynamodb_policy" {
  name = "lambda_dynamodb_policy"
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
          aws_dynamodb_table.posts_table.arn,
          "${aws_dynamodb_table.posts_table.arn}/*"
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_dynamodb_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_dynamodb_policy.arn
}