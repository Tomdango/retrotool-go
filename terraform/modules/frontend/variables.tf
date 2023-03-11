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
