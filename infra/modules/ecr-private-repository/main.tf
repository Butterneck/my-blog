terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}

data "aws_caller_identity" "current" {}
data "aws_partition" "current" {}


resource "aws_ecr_repository" "this" {

  name                 = var.name
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
