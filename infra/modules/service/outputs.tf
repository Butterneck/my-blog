output "api_id" {
  value = module.rest_api.api_id
}

output "api_stage_name" {
  value = module.rest_api.stage_name
}

output "api_region" {
  value = module.rest_api.region
}

output "ddb_cdc_bus_name" {
  value = length(module.dynamodb_table) > 0 ? module.dynamodb_table[0].cdc_bus_name : null
}

output "ddb_cdc_pipe_name" {
  value = length(module.dynamodb_table) > 0 ? module.dynamodb_table[0].cdc_pipe_name : null
}

output "ecr_repository_url" {
  value = module.ecr_repository.url
}