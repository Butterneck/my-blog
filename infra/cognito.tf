resource "aws_cognito_user_pool" "pool" {
    name = "blog"

    account_recovery_setting {
        recovery_mechanism {
            name     = "verified_email"
            priority = 1
        }        
    }

    alias_attributes = ["email", "preferred_username"]
    username_configuration {
        case_sensitive = false
    }

    mfa_configuration = "OPTIONAL"
    software_token_mfa_configuration {
      enabled = true
    }

    auto_verified_attributes = [ "email" ]

    deletion_protection = "ACTIVE"

}

resource "aws_cognito_user_pool_client" "blog-frontend" {
    name = "blog-frontend"
    user_pool_id = aws_cognito_user_pool.pool.id

    generate_secret = false

    allowed_oauth_flows_user_pool_client = true
    allowed_oauth_flows                  = ["code", "implicit"]
    allowed_oauth_scopes                 = ["email", "openid"]
    supported_identity_providers         = ["COGNITO"]
    callback_urls                        = ["https://example.com"]

    refresh_token_validity = 60
    access_token_validity = 1
    id_token_validity = 1

}
