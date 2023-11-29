provider "aws" {
  alias  = "euw1"
  region = "eu-west-1"
}

provider "aws" {
  alias  = "use1"
  region = "us-east-1"
}

####### TO BE MOVED TO A SEPARATE PROJECT #######

#############################################################
# DNS Zone that manages all the DNS records for the project #
#############################################################
module "dns_zone" {
  source = "./modules/hosted-zone"
  name   = local.blog_domain
}

##############################################################
# Permissions to allow API Gateway to create CloudWatch Logs #
##############################################################
resource "aws_api_gateway_account" "this" {
  cloudwatch_role_arn = aws_iam_role.api_gateway_cloudwatch_global.arn
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["apigateway.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "api_gateway_cloudwatch_global" {
  name               = "api-gateway-cloudwatch-global"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "api_gateway_cloudwatch_global" {
  statement {
    effect = "Allow"

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:DescribeLogGroups",
      "logs:DescribeLogStreams",
      "logs:PutLogEvents",
      "logs:GetLogEvents",
      "logs:FilterLogEvents",
    ]

    resources = ["*"]
  }
}
resource "aws_iam_role_policy" "api_gateway_cloudwatch_global" {
  name   = "default"
  role   = aws_iam_role.api_gateway_cloudwatch_global.id
  policy = data.aws_iam_policy_document.api_gateway_cloudwatch_global.json
}

####### END #######

################
# Blog Backend #
################
module "blog_backend" {
  source = "./modules/service"
  name   = "${local.name}-backend"
  dynamodb_config = {
    name = local.name
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
  openapi_file_path      = local.backend_openapi_file_path
  backend_image_uri      = var.backend_image_uri
  existing_user_pool_arn = module.user_pool.arn
}

#################
# Blog Frontend #
#################
module "blog_frontend" {
  source = "./modules/s3-bucket"
  name   = "${local.name}-frontend"
}

#############################################
# CDN for both the frontend and the backend #
#############################################
module "cdn" {
  source = "./modules/cdn"
  providers = {
    aws.main = aws.euw1
    aws.use1 = aws.use1
  }

  domain_names      = [local.blog_domain]
  domain_names_zone_name = local.blog_domain
  s3_origins = {
    "default" = {
      bucket_name = module.blog_frontend.name
      domain_name = module.blog_frontend.regional_domain_name
    }
  }
  apigw_origins = {
    backend = {
      rest_api_id = module.blog_backend.api_id
      stage_name  = module.blog_backend.api_stage_name
      region      = module.blog_backend.api_region
      path_pattern = "/api/*"
    }
  }

  # depends_on = [module.dns_zone, module.blog_frontend, module.blog_backend]
}

#####################################################
# Cognito User Pool used for backend authentication #
#####################################################
module "user_pool" {
  source = "./modules/cognito-user-pool"
  providers = {
    aws.main = aws.euw1
    aws.use1 = aws.use1
  }

  name   = local.name
  clients = {
    "frontend" = {
      callback_urls = [
        "https://${local.blog_domain}",
      ]
      logout_urls = [
        "https://${local.blog_domain}"
      ]
    }
  }
  custom_domain = "auth.${local.blog_domain}"
  custom_domain_zone_name = local.blog_domain

  admin_email = "pinton.filippo@protonmail.com"
  admin_username = "butterneck"
}

#####################
# Cache invalidator #
#####################
module "cache_invalidator_ecr_repository" {
  source = "./modules/ecr-private-repository"
  name   = "${local.name}-cache-invalidator"
}

module "cache_invalidator" {
  source         = "./modules/event-bridge-lambda"
  name           = "${local.name}-cache-invalidator"
  image_uri      = var.cache_invalidator_image_uri
  event_bus_name = module.blog_backend.ddb_cdc_bus_name
  event_pattern = jsonencode({
    source = [
      "Pipe ${module.blog_backend.ddb_cdc_pipe_name}"
    ]
  })
}
