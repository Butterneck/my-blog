module "lambda" {
    source    = "./../lambda"
    name      = var.name
    image_uri = var.image_uri
}

resource "aws_cloudwatch_event_rule" "rule" {
    name = var.name
    event_bus_name = var.event_bus_name
    event_pattern = var.event_pattern
}

resource "aws_cloudwatch_event_target" "target" {
    rule = aws_cloudwatch_event_rule.rule.name
    arn = module.lambda.arn
    event_bus_name = var.event_bus_name
}

resource "aws_lambda_permission" "permission" {
    statement_id = "AllowExecutionFromCloudWatch"
    action = "lambda:InvokeFunction"
    function_name = module.lambda.name
    principal = "events.amazonaws.com"
    source_arn = aws_cloudwatch_event_rule.rule.arn
}