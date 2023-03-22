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
    cognito_user_pool_web_client_id = dependency.cognito.outputs.cognito_user_pool_web_client_id
    cognito_user_pool_client_secret = dependency.cognito.outputs.cognito_user_pool_client_secret
    cognito_user_pool_endpoint = dependency.cognito.outputs.cognito_user_pool_endpoint
    private_subnets = dependency.networking.outputs.private_subnets

    db_host = dependency.database.outputs.host
    db_port = dependency.database.outputs.port
    db_username = dependency.database.outputs.username
    db_password = dependency.database.outputs.password
    db_name = dependency.database.outputs.db_name
}

dependency "build_infra" {
    config_path = "../build_infra"
}

dependency "cognito" {
    config_path = "../cognito"
}

dependency "networking" {
    config_path = "../networking"
}

dependency "database" {
    config_path = "../database"
}