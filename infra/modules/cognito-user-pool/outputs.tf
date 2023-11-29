output "arn" {
  value       = aws_cognito_user_pool.this.arn
  description = "ARN of the Cognito User Pool"
}
