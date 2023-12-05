terraform {
  required_version = "1.5.4"

  required_providers {
    aws = {
      source                = "hashicorp/aws"
      version               = ">= 5.26.0"
      configuration_aliases = [aws.main, aws.use1]
    }
  }
}

resource "aws_cognito_user_pool" "this" {
  name                = var.name
  deletion_protection = "INACTIVE"

  alias_attributes = ["email", "preferred_username"]
  username_configuration {
    case_sensitive = false
  }

  # Automatically send verification emails
  auto_verified_attributes = ["email"]

  # TODO: use SES instead
  email_configuration {
    email_sending_account = "COGNITO_DEFAULT"
  }

  # Disable self-signup
  admin_create_user_config {
    allow_admin_create_user_only = true
  }

  account_recovery_setting {
    recovery_mechanism {
      name     = "admin_only"
      priority = 1
    }
  }

  mfa_configuration = "ON"
  software_token_mfa_configuration {
    enabled = true
  }

}

resource "aws_cognito_user_pool_client" "blog-frontend" {
  for_each                             = var.clients
  name                                 = each.key
  user_pool_id                         = aws_cognito_user_pool.this.id
  allowed_oauth_flows                  = ["code"]
  callback_urls                        = each.value.callback_urls
  logout_urls                          = each.value.logout_urls
  allowed_oauth_flows_user_pool_client = true
  allowed_oauth_scopes                 = ["openid", "email", "profile"]
  supported_identity_providers         = ["COGNITO"]
}

module "custom_domain_certificate" {
  count            = var.custom_domain != null ? 1 : 0
  source           = "./../acm-certificate"
  domain_name      = var.custom_domain
  hosted_zone_name = var.custom_domain_zone_name

  providers = {
    aws = aws.use1
  }
}

resource "aws_cognito_user_pool_domain" "custom_domain" {
  count           = var.custom_domain != null ? 1 : 0
  domain          = var.custom_domain
  user_pool_id    = aws_cognito_user_pool.this.id
  certificate_arn = module.custom_domain_certificate[0].arn
}


# Admin user
resource "aws_cognito_user" "admin" {
  count        = var.admin_email != null || var.admin_username != null ? 1 : 0
  user_pool_id = aws_cognito_user_pool.this.id
  username     = coalesce(var.admin_username, var.admin_email)

  attributes = {
    email = var.admin_email
  }
}

data "aws_route53_zone" "custom_domain" {
  count = var.custom_domain != null ? 1 : 0
  name  = var.custom_domain_zone_name
}

resource "aws_route53_record" "auth_cognito_A" {
  count   = var.custom_domain != null ? 1 : 0
  name    = aws_cognito_user_pool_domain.custom_domain[0].domain
  type    = "A"
  zone_id = data.aws_route53_zone.custom_domain[0].zone_id
  alias {
    evaluate_target_health = false
    name                   = aws_cognito_user_pool_domain.custom_domain[0].cloudfront_distribution_arn
    # This zone_id is fixed
    zone_id = "Z2FDTNDATAQYW2"
  }
}
