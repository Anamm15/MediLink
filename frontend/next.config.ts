import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  images: {
    domains: [
      "source.unsplash.com",
      "images.unsplash.com",
      "res.cloudinary.com",
      "storage.googleapis.com",
      "i.pravatar.cc",
    ],
  },
};

export default nextConfig;
