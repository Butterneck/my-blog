resource "aws_s3_bucket" "blog_frontend" {
  bucket = "butterneck-blog-frontend"

  tags = {}
}

data "aws_iam_policy_document" "s3_policy" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.blog_frontend.arn}/*"]
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }
    condition {
      test     = "StringEquals"
      variable = "aws:SourceArn"
      values   = [aws_cloudfront_distribution.blog_frontend.arn]
    }
  }
}

resource "aws_s3_bucket_policy" "allow_cloudfront_read" {
  bucket = aws_s3_bucket.blog_frontend.id
  policy = data.aws_iam_policy_document.s3_policy.json
}

locals {
  frontend_origin_id = "s3Frontend"
  api_origin_id = "api"
}

resource "aws_cloudfront_origin_access_control" "blog_frontend" {
  name                              = "blog_frontend"
  description                       = "Access control for blog frontend"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_cloudfront_distribution" "blog_frontend" {
  origin {
    domain_name              = aws_s3_bucket.blog_frontend.bucket_regional_domain_name
    origin_id                = local.frontend_origin_id
    origin_access_control_id = aws_cloudfront_origin_access_control.blog_frontend.id
  }

  origin {
    domain_name = "${aws_api_gateway_rest_api.api.id}.execute-api.${var.aws_region}.amazonaws.com"
    origin_id   = local.api_origin_id
    origin_path = "/${aws_api_gateway_stage.prd.stage_name}"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1.2"]
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  comment             = "CDN for the blog frontend"
  default_root_object = "index.html"

  # logging_config {
  #   include_cookies = false
  #   bucket = "mylogs.s3.amazonaws.com"
  #   prefix = "myprefix"
  # }

  aliases = ["blog.butterneck.me"]

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = local.frontend_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }
    viewer_protocol_policy = "https-only"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  ordered_cache_behavior {
    path_pattern     = "/api/*"
    allowed_methods  = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "DELETE", "PATCH"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = local.api_origin_id

    forwarded_values {
      query_string = true

      cookies {
        forward = "none"
      }
    }
    viewer_protocol_policy = "https-only"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  price_class = "PriceClass_100"
  tags        = {}

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }


  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.blog_butterneck_me.arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2021"
  }

}

# DNS records
resource "aws_route53_record" "blog_butterneck_me" {
  name    = "blog.butterneck.me"
  type    = "A"
  zone_id = aws_route53_zone.blog_butterneck_me.zone_id

  alias {
    name                   = aws_cloudfront_distribution.blog_frontend.domain_name
    zone_id                = aws_cloudfront_distribution.blog_frontend.hosted_zone_id
    evaluate_target_health = true
  }
}

# ACM certificate

resource "aws_acm_certificate" "blog_butterneck_me" {
  provider          = aws.us-east-1
  domain_name       = "blog.butterneck.me"
  validation_method = "DNS"
}

resource "aws_route53_record" "blog_butterneck_me_validation_records" {
  for_each = {
    for dvo in aws_acm_certificate.blog_butterneck_me.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = aws_route53_zone.blog_butterneck_me.zone_id
}

resource "aws_acm_certificate_validation" "blog_butterneck_me" {
  provider                = aws.us-east-1
  certificate_arn         = aws_acm_certificate.blog_butterneck_me.arn
  validation_record_fqdns = [for record in aws_route53_record.blog_butterneck_me_validation_records : record.fqdn]
}