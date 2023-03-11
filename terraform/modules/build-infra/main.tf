/**
 * Build Artifact S3 Bucket
 */
resource "random_pet" "s3_bucket_seed" {}

module "artifact_s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"

  bucket = "${var.environment}-artifact-bucket-${random_pet.s3_bucket_seed.id}"
  acl    = "private"

  versioning = {
    enabled = false
  }
}


/**
 * ACM Certificate for domain
 */
data "aws_route53_zone" "root_hosted_zone" {
  name = var.root_domain_name
}

module "root_domain_certificate" {
  source = "terraform-aws-modules/acm/aws"

  domain_name = var.root_domain_name
  zone_id     = data.aws_route53_zone.root_hosted_zone.id

  subject_alternative_names = [
    "*.${var.root_domain_name}"
  ]

  wait_for_validation = true
}

/**
 * KMS Key for Encrypted Secrets
 */
resource "aws_kms_key" "sops_key" {
  description = "SOPS KMS Key for Encrypted Secrets"
}
