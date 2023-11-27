variable "name" {
  type        = string
  description = "The name of the record"
}

variable "hosted_zone_name" {
  type        = string
  description = "The hosted zone name"
}

variable "cloudfront_distribution_id" {
  type        = string
  description = "The cloudfront distribution id"
}
