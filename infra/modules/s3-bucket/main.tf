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


resource "aws_s3_bucket_notification" "bucket_notification" {
  count  = var.publish_events_on_eventbridge ? 1 : 0
  bucket = aws_s3_bucket.this.id

  eventbridge = true
}