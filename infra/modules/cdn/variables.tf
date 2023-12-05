variable "domain_names" {
  type        = list(string)
  description = "The domain names for the CloudFront distribution"
  default     = []
}

variable "domain_names_zone_name" {
  type        = string
  description = "The domain names zone name for the CloudFront distribution"
  default     = ""
}

variable "s3_origins" {
  type = map(object({
    bucket_name  = string
    domain_name  = string
    origin_path  = optional(string)
    path_pattern = optional(string)
  }))
  description = "The s3 origins for the CloudFront distribution. Must contain a `default` origin"
}

variable "apigw_origins" {
  type = map(object({
    rest_api_id  = string
    stage_name   = string
    path_pattern = string
    region       = string
  }))
  description = "The api gateway origins for the CloudFront distribution"
  default     = {}
}

variable "default_root_object" {
  type        = string
  description = "The default root object for the CloudFront distribution"
  default     = "index.html"
}

variable "cache_invalidation_events" {
  type = map(object({
    event_bus_name   = string,
    event_pattern    = string,
    invalidate_paths = optional(list(string))
  }))
  description = "The events that will trigger a cache invalidation"
  default     = {}
}
