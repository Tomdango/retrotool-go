generate "backend" {
  path      = "backend.tf"
  if_exists = "overwrite_terragrunt"
  contents = <<EOF
terraform {
  backend "s3" {
    bucket         = "retrotool-terraform-state"
    key            = "/prod/${path_relative_to_include()}/terraform.tfstate"
    region         = "eu-west-2"
    encrypt        = true
  }
}
EOF
}