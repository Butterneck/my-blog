output "arn" {
  value       = aws_cognito_user_pool.pool.arn
  description = "ARN of the Cognito User Pool"
}
