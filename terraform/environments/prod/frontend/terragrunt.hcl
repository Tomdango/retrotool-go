include {
    path = find_in_parent_folders()
}

terraform {
    source = "${path_relative_from_include()}//modules/frontend"
}

locals {
    secret_vars = yamldecode(sops_decrypt_file(find_in_parent_folders("secrets.yaml")))
}

inputs = {
    gh_token = local.secret_vars.gh_token,
    domain_name = "retrotool.tomjc.dev"
    environment = "prod"
}