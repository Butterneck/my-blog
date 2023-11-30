terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}

data "aws_region" "current" {}

resource "aws_api_gateway_rest_api" "this" {
  name = var.name

  body = templatefile(var.openapi_file_path, {
    integration_uri       = "arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/${var.backend_lambda_arn}/invocations"
    cognito_user_pool_arn = var.user_pool_arn
  })

  endpoint_configuration {
    types = ["REGIONAL"]
  }

  # disable_execute_api_endpoint = true
}

resource "aws_api_gateway_deployment" "this" {
  rest_api_id = aws_api_gateway_rest_api.this.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.this.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "main" {
  deployment_id = aws_api_gateway_deployment.this.id
  rest_api_id   = aws_api_gateway_rest_api.this.id
  stage_name    = "main"
}


# Enable logging and metrics for all methods
resource "aws_api_gateway_method_settings" "all" {
  rest_api_id = aws_api_gateway_rest_api.this.id
  stage_name  = aws_api_gateway_stage.main.stage_name
  method_path = "*/*"

  settings {
    metrics_enabled = true
    logging_level   = "INFO"
  }
}
