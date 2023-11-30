variable "name" {
  type        = string
  description = "Name of the Lambda function"
}

variable "image_uri" {
  type        = string
  description = "The ECR image URI containing the function's deployment package."
}


variable "events" {
  type = map(object({
    event_bus_name = string,
    event_pattern  = string
  }))
  description = "Map of events to subscribe to"
}
