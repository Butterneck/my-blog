terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.26.0"
    }
  }
}

resource "aws_route53_record" "this" {
  name    = var.name
  type    = "CNAME"
  zone_id = data.aws_route53_zone.selected.zone_id
  records = [var.value]
}

data "aws_route53_zone" "selected" {
  name         = var.hosted_zone_name
  private_zone = false
}
