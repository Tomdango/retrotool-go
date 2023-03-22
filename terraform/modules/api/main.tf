/**
 * API Lambda
 */
resource "aws_s3_object" "api_lambda_zip" {
  bucket = var.artifact_s3_bucket_id
  key    = "deploy/${var.environment}/${filemd5("${var.build_dir}/api.zip")}.zip"
  source = "${var.build_dir}/api.zip"
}


module "api_lambda_function" {
  source = "terraform-aws-modules/lambda/aws"

  function_name = "${var.environment}-api-lambda"
  description   = "Golang API Handler Lambda for ${var.environment}"
  runtime       = "go1.x"
  handler       = "main"

  create_package                          = false
  create_current_version_allowed_triggers = false

  s3_existing_package = {
    bucket = var.artifact_s3_bucket_id
    key    = aws_s3_object.api_lambda_zip.id
  }

  environment_variables = {
    COGNITO_USER_POOL_ID            = var.cognito_user_pool_id
    COGNITO_USER_POOL_CLIENT_ID     = var.cognito_user_pool_client_id
    COGNITO_USER_POOL_CLIENT_SECRET = var.cognito_user_pool_client_secret

    DB_HOST     = var.db_host
    DB_PORT     = var.db_port
    DB_USERNAME = var.db_username
    DB_PASSWORD = var.db_password
    DB_NAME     = var.db_name
  }

  allowed_triggers = {
    AllowExecutionFromAPIGateway = {
      service    = "apigateway"
      source_arn = "${module.api_gateway.apigatewayv2_api_execution_arn}/*/*"
    }
  }

  attach_policy_json = true
  policy_json        = <<-EOT
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": ["dynamodb:*"],
                "Resource": ["*"]
            }
        ]
    }
  EOT

  vpc_subnet_ids = var.private_subnets
}

/**
 * API Gateway
 */

module "api_gateway" {
  source = "terraform-aws-modules/apigateway-v2/aws"

  name          = "${var.environment}-api-gateway-http"
  description   = "HTTP API Gateway for ${var.environment}"
  protocol_type = "HTTP"

  cors_configuration = {
    allow_headers = ["content-type", "x-amz-date", "authorization", "x-api-key", "x-amz-security-token", "x-amz-user-agent"]
    allow_methods = ["*"]
    allow_origins = ["*"]
  }

  domain_name                  = "api.${var.root_domain_name}"
  domain_name_certificate_arn  = var.root_domain_certificate_arn
  disable_execute_api_endpoint = true

  integrations = {


    "POST /users/register" = {
      lambda_arn             = module.api_lambda_function.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds   = 12000
    }

    "POST /users/otp" = {
      lambda_arn             = module.api_lambda_function.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds   = 12000
    }

    "POST /users/login" = {
      lambda_arn             = module.api_lambda_function.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds   = 12000
    }

    "POST /teams/create" = {
      lambda_arn             = module.api_lambda_function.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds   = 12000
      authorization_type     = "JWT"
      authorizer_id          = aws_apigatewayv2_authorizer.api_auth.id
    }

    "ANY /v1/{proxy+}" = {
      lambda_arn             = module.api_lambda_function.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds   = 12000
    }
  }
}

resource "aws_apigatewayv2_authorizer" "api_auth" {
  api_id           = module.api_gateway.apigatewayv2_api_id
  authorizer_type  = "JWT"
  identity_sources = ["$request.header.Authorization"]
  name             = "api-cognito-auth"

  jwt_configuration {
    audience = [var.cognito_user_pool_web_client_id]
    issuer   = "https://${var.cognito_user_pool_endpoint}"
  }
}

data "aws_route53_zone" "this" {
  name = var.root_domain_name
}

resource "aws_route53_record" "api" {
  zone_id = data.aws_route53_zone.this.zone_id
  name    = "api"
  type    = "A"

  alias {
    name                   = module.api_gateway.apigatewayv2_domain_name_configuration[0].target_domain_name
    zone_id                = module.api_gateway.apigatewayv2_domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}
