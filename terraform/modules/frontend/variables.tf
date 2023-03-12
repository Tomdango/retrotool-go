variable "environment" {
  type        = string
  description = "Name of the environment that is being deployed to"
}

variable "gh_token" {
  type        = string
  sensitive   = true
  description = "Name of the environment that is being deployed to"
}

variable "domain_name" {
  type        = string
  description = "Domain name of the frontend"
}

variable "cognito_user_pool_id" {
  type = string
  description = "AWS Cognito User Pool ID"
}

variable "cognito_web_client_id" {
  type = string
  description = "AWS Cognito User Pool Web Client ID"
}