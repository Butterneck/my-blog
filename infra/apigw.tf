resource "aws_api_gateway_rest_api" "api" {
  name        = "blog-api"
  description = "API for blog"

  api_key_source = "HEADER"

  body = templatefile("${path.module}/../src/openapi.yaml", {
    backend_lambda_arn = aws_lambda_function.blog_backend.arn
    cognito_user_pool_arn = aws_cognito_user_pool.pool.arn
  })

  disable_execute_api_endpoint = false  # TODO: change to true when connecting to a custom domain
}


resource "aws_lambda_permission" "api_gateway" {
    statement_id  = "AllowExecutionFromAPIGateway"
    action        = "lambda:InvokeFunction"
    function_name = aws_lambda_function.blog_backend.function_name
    principal     = "apigateway.amazonaws.com"

    source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*"
}

resource "aws_api_gateway_deployment" "dev_api" {
    rest_api_id = aws_api_gateway_rest_api.api.id
}

resource "aws_api_gateway_stage" "dev" {
    stage_name    = "dev"
    rest_api_id   = aws_api_gateway_rest_api.api.id
    deployment_id = aws_api_gateway_deployment.dev_api.id
}