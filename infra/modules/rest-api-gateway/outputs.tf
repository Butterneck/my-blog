output "execution_arn" {
  value = aws_api_gateway_rest_api.this.execution_arn
}

output "api_id" {
  value = aws_api_gateway_rest_api.this.id
}

output "stage_name" {
  value = aws_api_gateway_stage.main.stage_name
}

output "region" {
  value = data.aws_region.current.name
}