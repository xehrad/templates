/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "standalone",
  reactStrictMode: true,
  typescript: {
    // LLM easy
    ignoreBuildErrors: true,
  },
  eslint: {
    // LLM easy
    ignoreDuringBuilds: true,
  },
};

export default nextConfig;
