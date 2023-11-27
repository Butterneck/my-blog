resource "aws_cognito_user_pool" "pool" {
  name                = var.user_pool_name
  alias_attributes    = ["email", "preferred_username"]
  deletion_protection = "ACTIVE"

  username_configuration {
    case_sensitive = false
  }
}

resource "aws_cognito_user_pool_client" "blog-frontend" {
  for_each     = toset(var.clients)
  name         = each.value
  user_pool_id = aws_cognito_user_pool.pool.id
}
