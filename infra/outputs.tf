output "backend_ecr_repo_url" {
  value = module.blog_backend.ecr_repository_url
}

output "frontend_bucket_name" {
  value = module.blog_frontend.name
}
