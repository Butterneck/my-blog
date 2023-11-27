output "name" {
  value = aws_lambda_function.this.function_name
}

output "arn" {
  value = aws_lambda_function.this.arn
}

output "role_arn" {
  value = aws_iam_role.lambda.arn
}

output "role_name" {
  value = aws_iam_role.lambda.name
}

output "cloudwatch_log_group_name" {
  value = aws_cloudwatch_log_group.lambda.name
}

output "cloudwatch_log_group_arn" {
  value = aws_cloudwatch_log_group.lambda.arn
}