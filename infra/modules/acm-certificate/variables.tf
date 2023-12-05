variable "domain_name" {
  type        = string
  description = "Domain name for the certificate"
}

variable "alternative_names" {
  type        = list(string)
  description = "Alternative names for the certificate"
  default     = []
}

variable "hosted_zone_name" {
  type        = string
  description = "Hosted zone name for the certificate"
}

variable "is_hosted_zone_private" {
  type        = bool
  description = "Is hosted zone private"
  default     = false
}

