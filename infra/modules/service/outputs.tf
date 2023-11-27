# api_id
# stage_name
# region

output "api_id" {
  value = aws_api_gateway_rest_api.this.id
}

output "stage_name" {
  value = aws_api_gateway_stage.this.stage_name
}

output "region" {
  value = data.aws_region.current.name
}