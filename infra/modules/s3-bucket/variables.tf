variable "name" {
  type        = string
  description = "The name of the bucket"
}

variable "publish_events_on_eventbridge" {
  type        = bool
  description = "Whether to publish events on EventBridge"
  default     = false
}
