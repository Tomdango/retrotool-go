variable "environment" {
  type        = string
  description = "Name of the environment that is being deployed to"
}

variable "vpc_cidr" {
  type        = string
  description = "The CIDR block of the VPC"
}
