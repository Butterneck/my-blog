terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}

data "aws_region" "current" {}

###################
# Lambda function #
###################

module "ecr_repository" {
  source = "./../ecr-private-repository"
  name   = var.name
}

resource "aws_lambda_permission" "api_gateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = module.lambda.name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${module.rest_api.execution_arn}/*/*"
}

module "lambda" {
  source             = "./../lambda"
  name               = var.name
  dynamodb_table_arn = module.dynamodb_table[0].arn
  dynamodb_table_name = module.dynamodb_table[0].name
  has_dynamodb_table = var.dynamodb_config != null
  image_uri          = var.backend_image_uri
}


##################
# DynamoDB table #
##################
module "dynamodb_table" {
  count                    = var.dynamodb_config != null ? 1 : 0
  source                   = "./../dynamodb-table"
  name                     = var.name
  attributes               = var.dynamodb_config.attributes
  hash_key                 = var.dynamodb_config.hash_key
  range_key                = var.dynamodb_config.range_key
  global_secondary_indexes = var.dynamodb_config.global_secondary_indexes
  expose_cdc_events        = var.dynamodb_config.expose_cdc_events
}


############
# Rest API #
############
module "rest_api" {
  source             = "./../rest-api-gateway"
  name               = var.name
  openapi_file_path  = var.openapi_file_path
  user_pool_arn      = var.existing_user_pool_arn
  backend_lambda_arn = module.lambda.arn
}
