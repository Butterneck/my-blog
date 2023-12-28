#!/bin/bash
set -xe

# Login to AWS ECR
aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 794148823865.dkr.ecr.eu-west-1.amazonaws.com

# Deploy the ecr repositories
terraform apply -target=module.blog_backend.module.ecr_repository -var "backend_image_uri=foo"

# Build and push the backend image
image_tag=$(date +%s)
ecr_repository=$(terraform output -raw backend_ecr_repo_url)
backend_image_uri=${ecr_repository}:${image_tag}
docker build --platform linux/arm64 -t ${backend_image_uri} -f ../blog-backend/Dockerfile ../blog-backend
docker push ${backend_image_uri}

# Deploy all the infratructure
terraform apply -var "backend_image_uri=${backend_image_uri}"

# # Build and deploy the frontend
# cd ../frontend
# pkgx npm install
# pkgx npm run build
# cd ../infra
# aws s3 sync ../frontend/build/ s3://$(terraform output -raw frontend_bucket_name) --delete