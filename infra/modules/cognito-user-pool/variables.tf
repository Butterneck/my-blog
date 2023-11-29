variable "name" {
  type        = string
  description = "Name of the Cognito User Pool"
}

variable "clients" {
  type        = map(object({
    callback_urls = list(string),
    logout_urls   = list(string),
  }))
  description = "List of clients config"
  default     = {}
}

variable "custom_domain" {
  type        = string
  description = "Domain name for the Cognito User Pool"
  default = null
}

variable "custom_domain_zone_name" {
  type        = string
  description = "Name of the r53 zone where custom_domain is hosted"
  default = null
}

variable "admin_email" {
  type        = string
  description = "Email of the admin user"
  default = null
}

variable "admin_username" {
  type        = string
  description = "Username of the admin user"
  default = null
}