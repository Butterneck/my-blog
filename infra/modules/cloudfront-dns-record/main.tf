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
  type    = "A"
  zone_id = aws_route53_zone.selected.zone_id

  alias {
    name                   = aws_cloudfront_distribution.selected.domain_name
    zone_id                = aws_cloudfront_distribution.selected.hosted_zone_id
    evaluate_target_health = true
  }
}

data "aws_route53_zone" "selected" {
  name         = var.hosted_zone_name
  private_zone = false
}

data "aws_cloudfront_distribution" "selected" {
  id = var.cloudfront_distribution_id
}
