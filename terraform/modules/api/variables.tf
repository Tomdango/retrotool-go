variable "build_dir" {
  type        = string
  description = "Absolute path of the build/ folder"
}

variable "environment" {
  type        = string
  description = "Name of the environment that is being deployed to"
}

variable "artifact_s3_bucket_id" {
  type        = string
  description = "The ID of the build artifact S3 bucket"
}

variable "root_domain_name" {
  type        = string
  description = "Environment root domain name"
}

variable "root_domain_certificate_arn" {
  type        = string
  description = "Environment root domain ACM certificate ARN"
}

variable "cognito_user_pool_id" {
  type        = string
  description = "Cognito User Pool ID"
}

variable "cognito_user_pool_web_client_id" {
  type        = string
  description = "Cognito User Pool Web Client ID"
}

variable "cognito_user_pool_endpoint" {
  type        = string
  description = "Cognito User Pool Endpoint"
}

variable "private_subnets" {
  type = list(string)
}