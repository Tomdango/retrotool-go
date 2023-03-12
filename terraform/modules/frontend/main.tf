
resource "aws_amplify_app" "frontend" {
  name         = "${var.environment}-retrotool-frontend"
  repository   = "https://github.com/Tomdango/retrotool-go"
  access_token = var.gh_token
  platform     = "WEB_COMPUTE"

  build_spec = <<-EOT
            version: 1
            applications:
              - frontend:
                  phases:
                    preBuild:
                      commands:
                        - yarn install
                    build:
                      commands:
                        - yarn run build
                  artifacts:
                    baseDirectory: .next
                    files:
                      - '**/*'
                  cache:
                    paths:
                      - node_modules/**/*
                appRoot: frontend
EOT

  environment_variables = {
    "AMPLIFY_MONOREPO_APP_ROOT"      = "frontend"
    "AMPLIFY_DIFF_DEPLOY"            = "false"
    "AWS_AUTH_REGION"                = "eu-west-2"
    "AWS_AUTH_USER_POOL_ID"          = var.cognito_user_pool_id
    "AWS_AUTH_USER_CLIENT_ID"        = var.cognito_web_client_id
    "AWS_AUTH_COOKIE_STORAGE_DOMAIN" = var.domain_name
  }

  custom_rule {
    source = "/<*>"
    status = "404"
    target = "/index.html"
  }
}

resource "aws_amplify_branch" "main" {
  app_id      = aws_amplify_app.frontend.id
  branch_name = "main"


  enable_basic_auth      = true
  basic_auth_credentials = base64encode("username:password")
}

resource "aws_amplify_domain_association" "domain" {
  app_id      = aws_amplify_app.frontend.id
  domain_name = var.domain_name

  sub_domain {
    branch_name = aws_amplify_branch.main.branch_name
    prefix      = ""
  }
}
