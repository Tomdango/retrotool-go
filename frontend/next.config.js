/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  env: {
    COGNITO_AUTH_REGION: process.env.COGNITO_AUTH_REGION,
    COGNITO_AUTH_USER_POOL_ID: process.env.COGNITO_AUTH_USER_POOL_ID,
    COGNITO_AUTH_USER_CLIENT_ID: process.env.COGNITO_AUTH_USER_CLIENT_ID,
    COGNITO_AUTH_COOKIE_STORAGE_DOMAIN:
      process.env.COGNITO_AUTH_COOKIE_STORAGE_DOMAIN,
  },
};

module.exports = nextConfig;
