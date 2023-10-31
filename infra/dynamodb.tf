resource "aws_dynamodb_table" "posts_table" {
  name                        = var.posts_table_name
  billing_mode                = "PAY_PER_REQUEST"
  hash_key                    = "id"
  range_key                   = "createdAt"
  deletion_protection_enabled = true

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "createdAt"
    type = "N"
  }

  attribute {
    name = "slug"
    type = "S"
  }

  global_secondary_index {
    name            = "slug-index"
    hash_key        = "slug"
    projection_type = "ALL"
  }
}

