variable "name" {
  type        = string
  description = "Name of the Service"
}

variable "openapi_file_path" {
  type        = string
  description = "Path to the OpenAPI file"
}

variable "existing_user_pool_arn" {
  type        = string
  description = "ARN of the existing Cognito User Pool to use. If not provided, and `requires_authentication=true` a new one will be created."
  default     = null
}

variable "dynamodb_config" {
  type = object({
    name                     = string
    attributes               = list(map(string))
    hash_key                 = string
    range_key                = string
    global_secondary_indexes = any
    expose_cdc_events        = bool
  })
  description = "Configuration for the DynamoDB table"
  default = null
}

variable "backend_image_uri" {
  type        = string
  description = "URI of the Docker image to use for the Lambda function"
}
