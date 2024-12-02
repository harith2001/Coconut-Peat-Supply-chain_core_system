#!/bin/bash

# Unzip the cutting.zip file from customPlugins directory
echo "Unzipping cutting.zip from customPlugins..."
if unzip customPlugins/cutting.zip -d customPlugins/unzipped_cutting; then
  echo "Unzip successful."
else
  echo "Failed to unzip cutting.zip. Exiting..."
  exit 1
fi

# Change directory to the unzipped folder
SPECIFIC_FOLDER="customPlugins/unzipped_cutting/cutting"
echo "Changing directory to $SPECIFIC_FOLDER"
cd "$SPECIFIC_FOLDER" || { echo "Failed to change directory to $SPECIFIC_FOLDER. Exiting..."; exit 1; }

# Check if the Dockerfile exists
DOCKERFILE="cutting.dockerfile"
if [ ! -f "$DOCKERFILE" ]; then
  echo "Dockerfile $DOCKERFILE not found. Exiting..."
  exit 1
fi

# Run go mod tidy
echo "Running go mod tidy..."
if go mod tidy; then
  echo "go mod tidy successful."
else
  echo "Failed to run go mod tidy. Exiting..."
  exit 1
fi

# Build the Docker image
echo "Building Docker image..."
if docker build -t cutting_plugin -f "$DOCKERFILE" .; then
  echo "Docker image build successful."
else
  echo "Failed to build Docker image. Exiting..."
  exit 1
fi

# Run the Docker container
echo "Running Docker container..."
if docker run -d -p 50053:50053 cutting_plugin; then
  echo "Docker container is up and running!"
else
  echo "Failed to run Docker container. Exiting..."
  exit 1
fi
