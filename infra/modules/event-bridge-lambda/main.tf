terraform {
  required_version = "1.5.4"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}

module "lambda" {
  source    = "./../lambda"
  name      = var.name
  image_uri = var.image_uri
  filename  = var.lambda_filename
  handler   = var.lambda_handler
  runtime   = var.lambda_runtime

  iam_assume_role_policy = var.lambda_iam_assume_role_policy
  iam_role_policies      = var.lambda_iam_role_policies
  environment_variables  = var.lambda_environment_variables
}

resource "aws_cloudwatch_event_rule" "rule" {
  for_each       = var.events
  name           = each.key
  event_bus_name = each.value.event_bus_name
  event_pattern  = each.value.event_pattern
}

resource "aws_cloudwatch_event_target" "target" {
  for_each       = var.events
  rule           = aws_cloudwatch_event_rule.rule[each.key].name
  arn            = module.lambda.arn
  event_bus_name = each.value.event_bus_name

  input = each.value.input
}

resource "aws_lambda_permission" "permission" {
  for_each      = var.events
  statement_id  = "AllowExecutionFromEventBridge-${each.key}"
  action        = "lambda:InvokeFunction"
  function_name = module.lambda.name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.rule[each.key].arn
}