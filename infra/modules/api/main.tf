###################
# Lambda function #
###################

module "ecr_repository" {
  source = "./modules/base/ecr-private-repository"
  name   = var.api_name
}

module "lambda" {
  source = "./modules/base/lambda"
  name   = var.api_name
  dynamodb_table_arn = module.dynamodb_table.arn
}


##################
# DynamoDB table #
##################
module "dynamodb_table" {
  source = "./modules/base/dynamodb-table"
  name   = var.api_name
  attributes = var.dynamodb_config.attributes
  hash_key = var.dynamodb_config.hash_key
  range_key = var.dynamodb_config.range_key
  global_secondary_indexes = var.dynamodb_config.global_secondary_indexes
  expose_cdc_events = var.dynamodb_config.expose_cdc_events
}


###############
# API Gateway #
###############

resource "aws_api_gateway_rest_api" "api" {
  name        = var.api_name

  body = templatefile(var.openapi_file_path, {
    integration_uri    = "arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/${module.lambda.arn}/invocations"
    cognito_user_pool_arn = coalesce(module.user_pool.arn, var.existing_user_pool_arn)
  })

  disable_execute_api_endpoint = true
}


resource "aws_lambda_permission" "api_gateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = module.lambda.name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*"
}


######################
# API Authentication #
######################

locals {
  create_user_pool = var.require_authentication && var.existing_user_pool_arn == ""
}

module "user_pool" {
  count = local.create_user_pool ? 1 : 0
  source = "./modules/base/cognito-user-pool"
  name   = var.api_name
}


