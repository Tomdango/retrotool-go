include {
    path = find_in_parent_folders()
}

terraform {
    source = "${path_relative_from_include()}//modules/build-infra"
}

inputs = {
    environment = "prod"
    root_domain_name = "retrotool.tomjc.dev"
}