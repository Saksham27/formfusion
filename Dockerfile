# Use Node.js LTS version as our base image
FROM node:20-slim

# Set working directory
WORKDIR /app

# Install basic development tools
RUN apt-get update && apt-get install -y \
    git \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Set environment variables
ENV NODE_ENV=development 