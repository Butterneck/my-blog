variable "api_name" {
    type = string
    description = "Name of the API"
    default = ""
}

variable "openapi_file_path" {
    type = string
    description = "Path to the OpenAPI file"
    default = ""
}

variable "existing_user_pool_arn" {
    type = string
    description = "ARN of the existing Cognito User Pool to use. If not provided, and `requires_authentication=true` a new one will be created."
    default = ""
}

variable "require_authentication" {
    type = bool
    description = "Whether the API requires authentication"
    default = false
}

variable "dynamodb_config" {
    type = object({
        name = string
        attributes = list(map(string))
        hash_key = string
        range_key = string
        global_secondary_indexes = any
        expose_cdc_events = bool
    })
    description = "Configuration for the DynamoDB table"
    default = {
        name = ""
        attributes = []
        hash_key = ""
        range_key = ""
        global_secondary_indexes = []
        expose_cdc_events = false
    }
}