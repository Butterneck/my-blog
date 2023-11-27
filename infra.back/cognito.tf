resource "aws_cognito_user_pool" "pool" {
  name = "blog"
  alias_attributes = ["email", "preferred_username"]
  deletion_protection = "ACTIVE"

  username_configuration {
    case_sensitive = false
  }
}

resource "aws_cognito_user_pool_client" "blog-frontend" {
  name         = "blog-frontend"
  user_pool_id = aws_cognito_user_pool.pool.id
}
