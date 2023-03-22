include {
    path = find_in_parent_folders()
}

terraform {
    source = "${path_relative_from_include()}//modules/database"
}

inputs = {
    environment = "prod"
    vpc_id = dependency.networking.outputs.vpc_id
    database_subnets = dependency.networking.outputs.database_subnets
    private_subnets_cidr_blocks = dependency.networking.outputs.private_subnets_cidr_blocks
}

dependency "networking" {
    config_path = "../networking"
}