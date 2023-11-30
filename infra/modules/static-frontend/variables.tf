variable "name" {
  type        = string
  description = "The name of the bucket"
}

variable "eventbridge_bus_name" {
  type        = string
  description = "The name of the event bus where the events will be published"
  default     = null
}
