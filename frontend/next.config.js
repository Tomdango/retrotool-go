/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  env: {
    AWS_AUTH_REGION: process.env.AWS_AUTH_REGION,
    AWS_AUTH_USER_POOL_ID: process.env.AWS_AUTH_USER_POOL_ID,
    AWS_AUTH_USER_CLIENT_ID: process.env.AWS_AUTH_USER_CLIENT_ID,
    AWS_AUTH_COOKIE_STORAGE_DOMAIN: process.env.AWS_AUTH_COOKIE_STORAGE_DOMAIN,
  },
};

module.exports = nextConfig;
