variable "name" {
  description = "The name of the API"
  type        = string

}

variable "openapi_file_path" {
  type        = string
  description = "Path to the OpenAPI file"
  default     = ""
}

variable "user_pool_arn" {
  type        = string
  description = "ARN of the existing Cognito User Pool to use for authentication (if needed)"
  default     = ""
}