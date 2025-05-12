#!/bin/bash

# Unzip the plugin.zip file from customPlugins directory

PLUGIN_NAME=$1

if [ -z "$PLUGIN_NAME" ]; then
  echo "Plugin name not provided. Exiting..."
  exit 1
fi

ZIP_FILE="plugins/${PLUGIN_NAME}.zip"
UNZIP_DIR="plugins/unzipped_${PLUGIN_NAME}"
PLUGIN_DIR="${UNZIP_DIR}/${PLUGIN_NAME}"
DOCKERFILE="${PLUGIN_NAME}.dockerfile"
KUBEFILE="${PLUGIN_NAME}-plugin.yaml"

# Check if the zip file exists
mkdir -p "$UNZIP_DIR"
if unzip -o "$ZIP_FILE" -d "$UNZIP_DIR"; then
  echo "Unzipped $ZIP_FILE successfully."
else
  echo "Failed to unzip $ZIP_FILE. Exiting..."
  exit 1
fi

# Navigate into plugin directory
if [ -d "$PLUGIN_DIR" ]; then
  cd "$PLUGIN_DIR"
else
  echo "Directory $PLUGIN_DIR does not exist. Exiting..."
  exit 1
fi

# Ensure files exist
[ -f "$DOCKERFILE" ] || { echo "$DOCKERFILE not found. Exiting..."; exit 1; }
[ -f "$KUBEFILE" ] || { echo "$KUBEFILE not found. Exiting..."; exit 1; }

# Run go mod tidy
echo "Running go mod tidy..."
if go mod tidy; then
  echo "go mod tidy successful."
else
  echo "Failed to run go mod tidy. Exiting..."
  exit 1
fi

# Ensure correct Docker Host for Rancher Desktop
unset DOCKER_HOST

# Docker Login
echo "Logging into Docker..."
if echo "harith2128" | docker login -u "harith2001" --password-stdin; then
  echo "Docker login successful."
else
  echo "Failed to login to Docker. Exiting..."
  exit 1
fi

# Build and push Docker image
IMAGE_NAME="harith2001/coconut-peat-supply-chain_core_system-${PLUGIN_NAME}:latest"
docker build -t "${PLUGIN_NAME}_plugin" -f "$DOCKERFILE" . || { echo "Docker build failed. Exiting..."; exit 1; }

# Push the Docker image
docker tag "${PLUGIN_NAME}_plugin" "$IMAGE_NAME"
docker push "$IMAGE_NAME" || { echo "Docker push failed. Exiting..."; exit 1; }

# Apply Kubernetes YAML
kubectl apply -f "$KUBEFILE" || { echo "Kube apply failed. Exiting..."; exit 1; }

echo "Plugin deployment complete for $PLUGIN_NAME."