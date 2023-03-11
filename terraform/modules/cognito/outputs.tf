
output "cognito_user_pool_id" {
  value = aws_cognito_user_pool.user_pool.id
}

output "cognito_user_pool_client_id" {
  value = aws_cognito_user_pool_client.client.id
}

output "cognito_user_pool_client_secret" {
  sensitive = true
  value     = aws_cognito_user_pool_client.client.client_secret
}