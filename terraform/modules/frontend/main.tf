
resource "aws_amplify_app" "frontend" {
  name         = "${var.environment}-retrotool-frontend"
  repository   = "https://github.com/Tomdango/retrotool-go"
  access_token = var.gh_token

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
            - yarn build
      artifacts:
        baseDirectory: .next
        files:
          - '**/*'
      cache:
        paths:
          - node_modules/**/*
    appRoot: frontend/
EOT

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
