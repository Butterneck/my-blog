variable "name" {
  description = "A unique name for your Lambda Function"
  type        = string
}

variable "cloudwatch_logs_retention_in_days" {
  description = "Specifies the number of days you want to retain log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653."
  type        = number
  default     = 14
}

variable "image_uri" {
  description = "The ECR image URI containing the function's deployment package."
  type        = string
  default     = null
}

variable "filename" {
  description = "The path to the function's deployment package within the local filesystem. If defined, The s3_bucket and s3_key variables will be ignored."
  type        = string
  default     = null
}

variable "handler" {
  description = "The function entrypoint in your code."
  type        = string
  default     = null
}

variable "runtime" {
  description = "The identifier of the function's runtime."
  type        = string
  default     = null
}

variable "image_config_entry_point" {
  description = "The ENTRYPOINT for the docker image"
  type        = list(string)
  default     = null
}
variable "image_config_command" {
  description = "The CMD for the docker image"
  type        = list(string)
  default     = null
}

variable "image_config_working_directory" {
  description = "The working directory for the docker image"
  type        = string
  default     = null
}

variable "dynamodb_table_arn" {
  description = "ARN of the DynamoDB table to which the Lambda function needs access"
  type        = string
  default     = null
}

variable "dynamodb_table_name" {
  description = "Name of the DynamoDB table to which the Lambda function needs access"
  type        = string
  default     = null
}

variable "has_dynamodb_table" {
  description = "Whether the Lambda function needs access to a DynamoDB table (needed explicitly because Terraform cannot detect if the DynamoDB table exists)"
  type        = bool
  default     = false
}

variable "iam_role_policies" {
  description = "A list of additional IAM policy ARNs to attach to the Lambda function's IAM role"
  type        = map(string)
  default     = {}
}

variable "iam_assume_role_policy" {
  description = "The IAM assume role policy document for the Lambda function's IAM role"
  type        = string
  default     = null
}

variable "environment_variables" {
  description = "A map that defines environment variables for the Lambda function"
  type        = map(string)
  default     = {}
}