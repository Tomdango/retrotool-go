output "artifact_s3_bucket_id" {
  value = module.artifact_s3_bucket.s3_bucket_id
}

output "acm_certificate_arn" {
  value = module.root_domain_certificate.acm_certificate_arn
}
