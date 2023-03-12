resource "random_string" "external_id" {
  length  = 12
  special = false
}

resource "aws_iam_role" "cognito_sms" {
  name = "${var.environment}-retrotool-sms-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "cognito-idp.amazonaws.com"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "sts:ExternalId": "${random_string.external_id.result}"
        }
      }
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "cognito_sms" {
  role = aws_iam_role.cognito_sms.name

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "sns:publish"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}
POLICY
}

resource "aws_cognito_user_pool" "user_pool" {
  name = "${var.environment}-retrotool-user-pool"

  auto_verified_attributes = ["phone_number"]
  sms_verification_message = "Your account activation code is {####}"

  admin_create_user_config {
    allow_admin_create_user_only = false
  }

  sms_configuration {
    sns_caller_arn = aws_iam_role.cognito_sms.arn
    external_id    = random_string.external_id.result
  }

  password_policy {
    minimum_length                   = 6
    require_lowercase                = true
    require_uppercase                = true
    require_numbers                  = true
    require_symbols                  = false
    temporary_password_validity_days = 2
  }

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_phone_number"
      priority = 1
    }
  }
}

resource "aws_cognito_user_pool_client" "client" {
  name = "${var.environment}-retrotool-cognito-client"

  user_pool_id            = aws_cognito_user_pool.user_pool.id
  generate_secret         = true
  enable_token_revocation = true

  refresh_token_validity        = 90
  prevent_user_existence_errors = "ENABLED"
  explicit_auth_flows = [
    "ALLOW_USER_SRP_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_ADMIN_USER_PASSWORD_AUTH"
  ]
}

resource "aws_cognito_user_pool_client" "web_client" {
  name = "${var.environment}-retrotool-cognito-web-client"

  user_pool_id            = aws_cognito_user_pool.user_pool.id
  generate_secret         = false
  enable_token_revocation = true

  refresh_token_validity        = 90
  prevent_user_existence_errors = "ENABLED"
  explicit_auth_flows = [
    "ALLOW_USER_SRP_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_ADMIN_USER_PASSWORD_AUTH"
  ]
}

resource "aws_cognito_user_pool_domain" "cognito_domain" {
  domain       = "${var.environment}-retrotool"
  user_pool_id = aws_cognito_user_pool.user_pool.id
}
