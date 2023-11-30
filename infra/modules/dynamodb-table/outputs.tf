output "arn" {
    value = aws_dynamodb_table.this.arn
}

output "name" {
    value = aws_dynamodb_table.this.name
}

output "cdc_bus_name" {
    value = local.eventbridge_bus_name
}

output "cdc_pipe_name" {
    value = aws_pipes_pipe.forward_cdc[0].name
}