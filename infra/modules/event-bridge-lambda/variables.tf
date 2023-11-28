variable "name" {
  type        = string
  description = "Name of the Lambda function"
}

variable "image_uri" {
  type        = string
  description = "The ECR image URI containing the function's deployment package."
}

variable "event_bus_name" {
  type        = string
  description = "Name of the EventBridge bus to use"
  default     = "default"
}

variable "event_pattern" {
  type        = string
  description = "The event pattern described a JSON object"
}
