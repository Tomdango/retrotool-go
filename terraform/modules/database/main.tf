data "aws_rds_engine_version" "postgresql" {
  engine  = "aurora-postgresql"
  version = "13.6"
}

resource "aws_db_parameter_group" "postgresql_13" {
  name        = "${var.environment}-aurora-db-postgres13-parameter-group"
  family      = "aurora-postgresql13"
  description = "${var.environment}-aurora-db-postgresql13-parameter-group"
}

module "aurora_postgresql_v2" {
  source  = "terraform-aws-modules/rds-aurora/aws"
  version = "~> 7.0"

  name              = "${var.environment}-postgresql"
  engine            = data.aws_rds_engine_version.postgresql.engine
  engine_mode       = "provisioned"
  engine_version    = data.aws_rds_engine_version.postgresql.version
  storage_encrypted = true
  database_name = "retrotool"

  vpc_id                = var.vpc_id
  subnets               = var.database_subnets
  create_security_group = true
  allowed_cidr_blocks   = var.private_subnets_cidr_blocks

  monitoring_interval = 60

  apply_immediately   = true
  skip_final_snapshot = true

  db_parameter_group_name = aws_db_parameter_group.postgresql_13.id

  serverlessv2_scaling_configuration = {
    min_capacity = 1
    max_capacity = 5
  }

  instance_class = "db.serverless"
  instances = {
    one = {}
  }
}

resource "aws_ssm_parameter" "db_name" {
  name  = "/${var.environment}/db/name"
  type  = "SecureString"
  value = module.aurora_postgresql_v2.cluster_database_name
}

resource "aws_ssm_parameter" "db_endpoint" {
  name  = "/${var.environment}/db/endpoint"
  type  = "SecureString"
  value = module.aurora_postgresql_v2.cluster_endpoint
}

resource "aws_ssm_parameter" "db_port" {
  name  = "/${var.environment}/db/port"
  type  = "SecureString"
  value = module.aurora_postgresql_v2.cluster_port
}

resource "aws_ssm_parameter" "db_username" {
  name  = "/${var.environment}/db/username"
  type  = "SecureString"
  value = module.aurora_postgresql_v2.cluster_master_username
}

resource "aws_ssm_parameter" "db_password" {
  name  = "/${var.environment}/db/password"
  type  = "SecureString"
  value = module.aurora_postgresql_v2.cluster_master_password
}