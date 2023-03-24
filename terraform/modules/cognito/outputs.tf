output "cognito_user_pool_id" {
  value = aws_cognito_user_pool.user_pool.id
}

output "cognito_user_pool_endpoint" {
  value = aws_cognito_user_pool.user_pool.endpoint
}

output "cognito_user_pool_web_client_id" {
  value = aws_cognito_user_pool_client.web_client.id
}