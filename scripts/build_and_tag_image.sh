#!/bin/bash

# Replace the image name and tag with version + commit SHA
REPOSITORY_URL="ledger"

# Get the script directory
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

# Read the version number from the .version file in the script directory
VERSION=$(cat "${SCRIPT_DIR}/../.version")

TAG="${VERSION}.$(git rev-parse --short HEAD)"

docker build --platform linux/arm64 -t "${REPOSITORY_URL}:${TAG}" -t "${REPOSITORY_URL}:latest" .

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Docker image built successfully: ARM64"
else
  echo "Error: Docker image build failed - ARM64"
  exit 1
fi

echo "Docker image tagged as:"
echo "  ${REPOSITORY_URL}:${TAG}"
echo "  ${REPOSITORY_URL}:latest"
