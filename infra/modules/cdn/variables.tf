variable "domain_names" {
  type        = list(string)
  description = "The domain names for the CloudFront distribution"
  default     = []
}

variable "domain_names_zone" {
  type        = string
  description = "The domain names zone for the CloudFront distribution"
  default     = ""
}

variable "s3_origins" {
  type = list(map(object({
    id          = string
    bucket_name = string
    origin_path = string
  })))
  description = "The s3 origins for the CloudFront distribution. The first origin is the default origin."
  default     = []
}

variable "apigw_origins" {
  type = list(map(object({
    id          = string
    rest_api_id = string
    stage_name  = string
    region      = string
  })))
  description = "The api gateway origins for the CloudFront distribution"
  default     = []
}

variable "default_root_object" {
  type        = string
  description = "The default root object for the CloudFront distribution"
  default     = "index.html"
}