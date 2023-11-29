# Build and deploy the frontend
cd ../frontend
pkgx npm install
pkgx npm run build
cd ../infra
aws s3 sync ../frontend/build/ s3://$(terraform output -raw frontend_bucket_name) --delete