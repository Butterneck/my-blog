variable "backend_image_uri" {
  type        = string
  description = "URI of the container image to deploy"
  default     = ""
}

variable "backend_openapi_file_path" {
  type        = string
  description = "Path to the OpenAPI file"
  default     = ""
}