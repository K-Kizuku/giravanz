resource "aws_cloudfront_distribution" "static-www" {
    origin {
        domain_name = aws_s3_bucket.frontend_bucket.bucket_regional_domain_name
        origin_id = aws_s3_bucket.frontend_bucket.id
        s3_origin_config {
          origin_access_identity = aws_cloudfront_origin_access_identity.static-www.cloudfront_access_identity_path
        }
    }

    enabled =  true

    default_root_object = "index.html"

    default_cache_behavior {
        allowed_methods = [ "GET", "HEAD" ]
        cached_methods = [ "GET", "HEAD" ]
        target_origin_id = aws_s3_bucket.frontend_bucket.id
        
        forwarded_values {
            query_string = false

            cookies {
              forward = "none"
            }
        }

        viewer_protocol_policy = "redirect-to-https"
        min_ttl = 0
        default_ttl = 3600
        max_ttl = 86400
    }
    aliases = [ "kizuku-hackathon.work" ]


    restrictions {
      geo_restriction {
          restriction_type = "whitelist"
          locations = [ "JP" ]
      }
    }
  viewer_certificate {
    acm_certificate_arn = "arn:aws:acm:us-east-1:290517700846:certificate/5a667b8a-ef44-4e0d-a72e-68a4fffa6238"
    ssl_support_method  = "sni-only"
    minimum_protocol_version = "TLSv1.2_2019"
  }
    
}

resource "aws_cloudfront_origin_access_identity" "static-www" {}