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
    range_key                = optional(string)
    global_secondary_indexes = optional(any)
    expose_cdc_events        = bool
    eventbridge_bus_name     = optional(string)
  })
  description = "Configuration for the DynamoDB table"
  default     = null
}

variable "backend_image_uri" {
  type        = string
  description = "The ECR image URI containing the backend's deployment package."
}

variable "iam_role_policies" {
  description = "A list of additional JSON formatted IAM policies to attach to the Lambda function's IAM role"
  type        = map(string)
  default     = {}
}