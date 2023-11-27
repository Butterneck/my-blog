terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}

resource "aws_s3_bucket" "this" {
  bucket = var.name
}
