output "host" {
  value = module.aurora_postgresql_v2.cluster_endpoint
}

output "port" {
  value = module.aurora_postgresql_v2.cluster_port
}

output "username" {
  sensitive = true
  value     = module.aurora_postgresql_v2.cluster_master_username
}

output "password" {
  sensitive = true
  value     = module.aurora_postgresql_v2.cluster_master_password
}

output "db_name" {
  value = module.aurora_postgresql_v2.cluster_database_name
}
