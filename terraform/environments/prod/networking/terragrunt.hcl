include {
    path = find_in_parent_folders()
}

terraform {
    source = "${path_relative_from_include()}//modules/networking"
}

inputs = {
    environment = "prod"
    vpc_cidr = "10.0.0.0/18"
}