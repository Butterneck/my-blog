###########################
# CloudFront Distribution #
###########################
resource "aws_cloudfront_distribution" "this" {

  dynamic "origin" {
    for_each = var.s3_origins
    content {
      domain_name = origin.value.domain_name
      origin_id   = origin.value.origin_id
      #   origin_access_identity = origin.value.origin_access_identity
    }
  }

  dynamic "origin" {
    for_each = var.apigw_origins
    content {
      domain_name = origin.value.domain_name
      origin_id   = origin.value.origin_id
      origin_path = origin.value.path
      custom_origin_config {
        http_port              = 80
        https_port             = 443
        origin_protocol_policy = "https-only"
        origin_ssl_protocols   = ["TLSv1.2"]
      }
    }
  }

  enabled         = true
  is_ipv6_enabled = true
  comment         = ""

  default_cache_behavior {
    allowed_methods        = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = ""
    viewer_protocol_policy = "https-only"

  }

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.blog_butterneck_me.arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2021"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
}


#######################################
# Origin access control for S3 origin #
#######################################
resource "aws_cloudfront_origin_access_control" "foo" {
  name                              = "${each.value.origin_id}-origin-access-control"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}
