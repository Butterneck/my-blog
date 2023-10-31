output "backend_lambda_name" {
  description = "Backend lambda function name"
  value       = aws_lambda_function.blog_backend.function_name
}

# output "base_url" {
#   description = "Base URL for API Gateway stage."

#   value = 
# }
