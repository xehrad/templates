import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,

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
