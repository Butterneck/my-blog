terraform {
  required_version = "1.5.4"

  required_providers {
    aws = {
      source                = "hashicorp/aws"
      version               = ">= 5.26.0"
      configuration_aliases = [aws.main, aws.use1]
    }

    archive = {
      source  = "hashicorp/archive"
      version = "2.4.0"
    }
  }
}

resource "aws_cloudfront_distribution" "this" {

  aliases = var.domain_names

  dynamic "origin" {
    for_each = var.s3_origins
    content {
      domain_name              = origin.value.domain_name
      origin_id                = origin.key
      origin_access_control_id = aws_cloudfront_origin_access_control.blog_frontend.id
      origin_path              = origin.value.origin_path
    }
  }

  dynamic "origin" {
    for_each = var.apigw_origins
    content {
      domain_name = "${origin.value.rest_api_id}.execute-api.${origin.value.region}.amazonaws.com"
      origin_id   = origin.key
      origin_path = "/${origin.value.stage_name}"

      custom_origin_config {
        http_port              = 80
        https_port             = 443
        origin_protocol_policy = "https-only"
        origin_ssl_protocols   = ["TLSv1.2"]
      }
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  price_class         = "PriceClass_100"
  default_root_object = var.default_root_object

  # Default cache behavior optimized for s3 origin
  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "DELETE", "PATCH"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = "default"

    # Managed-CachingOptimized
    cache_policy_id = "658327ea-f89d-4fab-a63d-7e88639e58f6"

    viewer_protocol_policy = "https-only"
  }

  dynamic "ordered_cache_behavior" {
    for_each = var.apigw_origins
    content {
      path_pattern     = ordered_cache_behavior.value.path_pattern
      allowed_methods  = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "DELETE", "PATCH"]
      cached_methods   = ["GET", "HEAD", "OPTIONS"]
      target_origin_id = ordered_cache_behavior.key

      #   # Managed-CORS-with-preflight-and-SecurityHeadersPolicy
      #   response_headers_policy_id = "eaab4381-ed33-4a86-88ca-d9558dc6cd63"

      #   Managed-SecurityHeadersPolicy
      response_headers_policy_id = "67f7725c-6f97-4210-82d7-5512b31e9d03"

      # Managed-AllViewerExceptHostHeader
      origin_request_policy_id = "b689b0a8-53d0-40ab-baf2-68738e2966ac"

      # Managed-CachingDisabled
      cache_policy_id = "4135ea2d-6df8-44a3-9df3-4b5a84be39ad"

      viewer_protocol_policy = "https-only"
      min_ttl                = 0
      default_ttl            = 0
      max_ttl                = 0
    }
  }

  dynamic "ordered_cache_behavior" {
    for_each = {
      for id, s3_origin in var.s3_origins : id => s3_origin if id != "default"
    }
    content {
      allowed_methods  = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "DELETE", "PATCH"]
      cached_methods   = ["GET", "HEAD", "OPTIONS"]
      target_origin_id = ordered_cache_behavior.key
      path_pattern     = ordered_cache_behavior.value.path_pattern

      # Managed-CachingOptimized
      cache_policy_id = "658327ea-f89d-4fab-a63d-7e88639e58f6"

      viewer_protocol_policy = "https-only"
      min_ttl                = 0
      # default_ttl            = 3600
      # max_ttl                = 86400
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn      = module.acm_certificate.arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2021"
  }

  custom_error_response {
    error_caching_min_ttl = 10
    error_code            = 403
    response_code         = 200
    response_page_path    = "/index.html"
  }
}

resource "aws_cloudfront_origin_access_control" "blog_frontend" {
  name                              = var.domain_names[0]
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

################
# Certificates #
################
module "acm_certificate" {
  source = "./../acm-certificate"

  domain_name            = var.domain_names[0]
  alternative_names      = length(var.domain_names) > 1 ? slice(var.domain_names, 1, length(var.domain_names) - 1) : []
  hosted_zone_name       = var.domain_names_zone_name
  is_hosted_zone_private = false

  providers = {
    aws = aws.use1
  }
}

###############
# DNS records #
###############
module "r53_records" {
  for_each = toset(var.domain_names)
  source   = "./../cloudfront-dns-record"

  name             = each.value
  hosted_zone_name = var.domain_names_zone_name

  cloudfront_distribution_id = resource.aws_cloudfront_distribution.this.id
}

##############
# S3 origins #
##############
data "aws_s3_bucket" "s3_origin" {
  for_each = var.s3_origins
  bucket   = each.value.bucket_name
}

##################
# S3 permissions #
##################
resource "aws_s3_bucket_policy" "allow_cloudfront_read" {
  for_each = var.s3_origins
  bucket   = data.aws_s3_bucket.s3_origin[each.key].id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action   = "s3:GetObject"
        Resource = "${data.aws_s3_bucket.s3_origin[each.key].arn}/*"
        Condition = {
          StringEquals = {
            "aws:SourceArn" = aws_cloudfront_distribution.this.arn
          }
        }
      }
    ]
  })
}

locals {
  _cache_invalidation_events = { for event_name, event in var.cache_invalidation_events : event_name => {
    event_bus_name = event.event_bus_name
    event_pattern  = event.event_pattern
    input = jsonencode({
      invalidate_paths = event.invalidate_paths != null ? event.invalidate_paths : ["/*"]
    })
  } }
}

module "cache_invalidator" {
  source          = "./../event-bridge-lambda"
  name            = "cloudfront-${aws_cloudfront_distribution.this.id}-cache-invalidator"
  lambda_filename = data.archive_file.cache_invalidator_zip.output_path
  lambda_handler  = "index.handler"
  lambda_runtime  = "python3.11"

  events = local._cache_invalidation_events

  lambda_environment_variables = {
    CLOUDFRONT_DISTRIBUTION_ID = aws_cloudfront_distribution.this.id
  }

  lambda_iam_role_policies = {
    "cloudfront_invalidate_cache" = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Action = [
            "cloudfront:CreateInvalidation",
          ]
          Effect   = "Allow"
          Resource = aws_cloudfront_distribution.this.arn
        },
      ]
    })
  }
}

data "archive_file" "cache_invalidator_zip" {
  type        = "zip"
  source_file = "${path.module}/assets/invalidate-cloudfront-cache/index.py"
  output_path = "${path.module}/assets/invalidate-cloudfront-cache.zip"
}
