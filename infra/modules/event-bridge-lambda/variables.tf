variable "name" {
  type        = string
  description = "Name of the Lambda function"
}

variable "image_uri" {
  type        = string
  description = "The ECR image URI containing the function's deployment package."
  default     = null
}


variable "events" {
  type = map(object({
    event_bus_name = string,
    event_pattern  = string,
    input          = optional(string)
  }))
  description = "Map of events to subscribe to"
}

variable "lambda_filename" {
  description = "The path to the function's deployment package within the local filesystem. If defined, The s3_bucket and s3_key variables will be ignored."
  type        = string
  default     = null
}

variable "lambda_handler" {
  description = "The function entrypoint in your code."
  type        = string
  default     = null
}

variable "lambda_runtime" {
  description = "The identifier of the function's runtime."
  type        = string
  default     = null
}


variable "lambda_iam_role_policies" {
  description = "A list of additional IAM policy ARNs to attach to the Lambda function's IAM role"
  type        = map(string)
  default     = {}
}

variable "lambda_iam_assume_role_policy" {
  description = "The IAM assume role policy document for the Lambda function's IAM role"
  type        = string
  default     = null
}

variable "lambda_environment_variables" {
  description = "A map that defines environment variables for the Lambda function"
  type        = map(string)
  default     = {}
}
