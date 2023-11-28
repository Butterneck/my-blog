output "arn" {
    value = aws_dynamodb_table.this.arn
}

output "cdc_bus_name" {
    value = aws_cloudwatch_event_bus.cdc_bus[0].name
}

output "cdc_pipe_name" {
    value = aws_pipes_pipe.forward_cdc[0].name
}