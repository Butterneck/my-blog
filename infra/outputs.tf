output "backend_ecr_repo_url" {
  value = module.blog_backend.ecr_repository_url
}

output "cache_invalidator_repo_url" {
  value = module.cache_invalidator_ecr_repository.url
}

output "frontend_bucket_name" {
  value = module.blog_frontend.name
}
