variable "name" {
  type        = string
  description = "Name of the Lambda function"
}

variable "sqs_queue_name" {
  type        = string
  description = "Name of the SQS queue to which the Lambda function needs access"
  default     = null
}

variable "sqs_queue_policies" {
  type = map(object({
    actions          = list(string),
    servicePrincipal = string,
    sourceArn        = string,
  }))
  description = "Map of SQS queue policies to apply"
  default     = {}
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


variable "lambda_image_uri" {
  description = "The ECR image URI containing the function's deployment package."
  type        = string
  default     = null
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

variable "batch_size" {
  description = "The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function."
  type        = number
  default     = 10000
}

variable "max_batching_window_in_seconds" {
  description = "The maximum amount of time to gather records before invoking the function, in seconds."
  type        = number
  default     = 60
}

variable "max_concurrency" {
  description = "The maximum number of concurrent invocations reserved for the function."
  type        = number
  default     = 2
}

variable "lambda_environment_variables" {
  description = "A map that defines environment variables for the Lambda function"
  type        = map(string)
  default     = {}
}
