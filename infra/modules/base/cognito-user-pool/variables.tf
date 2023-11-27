variable "user_pool_name" {
    type = string
    description = "Name of the Cognito User Pool"
    default = ""
}

variable "clients" {
    type = list(string)
    description = "List of client names"
    default = []
}