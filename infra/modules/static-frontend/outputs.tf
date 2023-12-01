output "name" {
  value = module.s3_bucket.name
}

output "domain_name" {
  value = module.s3_bucket.domain_name
}

output "regional_domain_name" {
  value = module.s3_bucket.regional_domain_name
}

output "deploy_events_bus_name" {
  value = local.eventbridge_bus_name
}
