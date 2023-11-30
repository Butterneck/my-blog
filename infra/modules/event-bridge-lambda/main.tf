terraform {
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
}

resource "aws_lambda_permission" "permission" {
  for_each      = var.events
  statement_id  = "AllowExecutionFromEventBridge-${each.key}"
  action        = "lambda:InvokeFunction"
  function_name = module.lambda.name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.rule[each.key].arn
}