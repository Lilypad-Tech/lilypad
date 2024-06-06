# Using Ubuntu as the base image
FROM ubuntu:latest

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    git \
    curl \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Install Lilypad
COPY ./lilypad /usr/local/bin/lilypad
RUN chmod +x /usr/local/bin/lilypad
ENV PATH="/usr/local/bin:${PATH}"