variable "backend_image_uri" {
  type        = string
  description = "URI of the container image to deploy for blog backend"
}

variable "cache_invalidator_image_uri" {
  type        = string
  description = "URI of the container image to deploy for cache invalidator"
}
