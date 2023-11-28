provider "aws" {
  alias = "euw1"
  region = "eu-west-1"
}

provider "aws" {
  alias  = "use1"
  region = "us-east-1"
}

module "blog_backend" {
  source = "./modules/service"
  name   = var.name
  dynamodb_config = {
    name = var.name
    attributes = [
      {
        name = "id"
        type = "S"
      },
      {
        name = "createdAt"
        type = "S"
      },
      {
        name = "slug"
        type = "S"
      }
    ]
    hash_key  = "id"
    range_key = "createdAt"
    global_secondary_indexes = [
      {
        name            = "slug-index"
        hash_key        = "slug"
        projection_type = "ALL"
      }
    ]
    expose_cdc_events = true
  }
  openapi_file_path = var.backend_openapi_file_path
  backend_image_uri = var.backend_image_uri
  existing_user_pool_arn = module.user_pool.arn
}

module "blog_frontend" {
  source = "./modules/s3-bucket"
  name   = var.name
}

module "dns_zone" {
  source = "./modules/hosted-zone"
  name   = local.blog_domain
}

module "cdn" {
  source            = "./modules/cdn"
  providers = {
    aws.main = aws.euw1
    aws.use1 = aws.use1
  }

  domain_names      = [local.blog_domain]
  domain_names_zone = local.blog_domain
  s3_origins = {
    "default" = {
      bucket_name = module.blog_frontend.name
      domain_name = module.blog_frontend.regional_domain_name
      origin_path = "/"
    }
  }
  apigw_origins = {
    backend = {
      rest_api_id = module.blog_backend.api_id
      stage_name  = module.blog_backend.api_stage_name
      region      = module.blog_backend.api_region
    }
  }

  depends_on = [ module.dns_zone, module.blog_frontend, module.blog_backend ]
}

######################
# Cognito User Pool #
######################
module "user_pool" {
  source = "./modules/cognito-user-pool"
  name   = var.name
}
