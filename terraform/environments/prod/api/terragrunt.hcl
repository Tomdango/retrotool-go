include {
    path = find_in_parent_folders()
}

terraform {
    source = "${path_relative_from_include()}//modules/api"
}

inputs = {
    build_dir = "${get_repo_root()}/build"
    environment = "prod"
    root_domain_name = "retrotool.tomjc.dev"
    root_domain_certificate_arn = dependency.build_infra.outputs.acm_certificate_arn
    artifact_s3_bucket_id = dependency.build_infra.outputs.artifact_s3_bucket_id
    cognito_user_pool_id = dependency.cognito.outputs.cognito_user_pool_id
    cognito_user_pool_client_id = dependency.cognito.outputs.cognito_user_pool_client_id   
    cognito_user_pool_client_secret = dependency.cognito.outputs.cognito_user_pool_client_secret
}

dependency "build_infra" {
    config_path = "../build_infra"
}

dependency "cognito" {
    config_path = "../cognito"
}