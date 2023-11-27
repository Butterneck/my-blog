variable "function_name" {
  description = "A unique name for your Lambda Function"
  type        = string
  default     = ""
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

variable "image_config_entry_point" {
  description = "The ENTRYPOINT for the docker image"
  type        = list(string)
  default     = []

}
variable "image_config_command" {
  description = "The CMD for the docker image"
  type        = list(string)
  default     = []
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