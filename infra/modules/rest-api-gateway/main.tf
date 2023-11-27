resource "aws_api_gateway_rest_api" "api" {
  name = var.name

  body = templatefile(var.openapi_file_path, {
    integration_uri       = "arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/${module.lambda.arn}/invocations"
    cognito_user_pool_arn = coalesce(module.user_pool.arn, var.user_pool_arn)
  })

  disable_execute_api_endpoint = true
}

resource "aws_api_gateway_deployment" "main" {
  rest_api_id = aws_api_gateway_rest_api.api.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.api.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "main" {
  deployment_id = aws_api_gateway_deployment.api.id
  rest_api_id   = aws_api_gateway_rest_api.api.id
  stage_name    = "main"
}
