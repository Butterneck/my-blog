variable "backend_image_uri" {
  type        = string
  description = "URI of the container image to deploy"
  default     = "" # TODO: remove default value
}

variable "backend_openapi_file_path" {
  type        = string
  description = "Path to the OpenAPI file"
  default     = "../src/blog-backend/openapi.yaml" # TODO: remove default value
}

variable "name" {
  type        = string
  description = "Name of the Service"
  default     = "butterneck-me-blog"
}