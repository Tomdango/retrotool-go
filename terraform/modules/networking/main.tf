module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = "${var.environment}-vpc"
  cidr = var.vpc_cidr

  enable_dns_support   = true
  enable_dns_hostnames = true

  azs              = ["eu-west-2a", "eu-west-2b"]
  public_subnets   = [cidrsubnet(var.vpc_cidr, 6, 0), cidrsubnet(var.vpc_cidr, 6, 1)]
  private_subnets  = [cidrsubnet(var.vpc_cidr, 6, 2), cidrsubnet(var.vpc_cidr, 6, 3)]
  database_subnets = [cidrsubnet(var.vpc_cidr, 6, 4), cidrsubnet(var.vpc_cidr, 6, 5)]
}
