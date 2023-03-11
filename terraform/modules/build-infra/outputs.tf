output "artifact_s3_bucket_id" {
  value = module.artifact_s3_bucket.s3_bucket_id
}

output "acm_certificate_arn" {
  value = module.root_domain_certificate.acm_certificate_arn
}

output "sops_kms_key_arn" {
  value = aws_kms_key.sops_key.arn
}
