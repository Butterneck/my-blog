variable "apigw_origins" {
  description = "List of API Gateway origins"
  type        = list(object({
    domain_name = string
    origin_id   = string
    path        = string
  }))
    default     = []
}

variable "s3_origins" {
  description = "List of S3 origins"
  type        = list(map(string))
  default     = []
}