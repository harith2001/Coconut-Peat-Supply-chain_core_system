#!/bin/bash

# Unzip the plugin.zip file from customPlugins directory
echo "Unzipping washing.zip from Plugins..."
rm -rf plugins/unzipped_washing
mkdir -p plugins/unzipped_washing

if unzip -o plugins/washing.zip -d plugins/unzipped_washing; then
  echo "Unzip successful."
else
  echo "Failed to unzip washing.zip. Exiting..."
  exit 1
fi

# Change directory to the unzipped folder
SPECIFIC_FOLDER="plugins/unzipped_washing/washing"
if [ -d "$SPECIFIC_FOLDER" ]; then
  echo "Changing directory to $SPECIFIC_FOLDER"
  cd "$SPECIFIC_FOLDER"
else
  echo "Directory $SPECIFIC_FOLDER does not exist. Exiting..."
  exit 1
fi

# Check if the Dockerfile exists
DOCKERFILE="washing.dockerfile"
if [ ! -f "$DOCKERFILE" ]; then
  echo "Dockerfile $DOCKERFILE not found. Exiting..."
  exit 1
fi

# Check if the kube yaml file exists
KUBEFILE="washing-plugin.yaml"
if [ ! -f "$KUBEFILE" ]; then
  echo "Kube yaml file $KUBEFILE not found. Exiting..."
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


# Build the Docker image
echo "Building Docker image..."
if docker build -t washing_plugin -f "$DOCKERFILE" .; then
  echo $DOCKERFILE
  echo "Docker image build successful."
else
  echo "Failed to build Docker image. Exiting..."
  exit 1
fi

# Push the Docker image
echo "Pushing Docker image as latest..."
IMAGE_TAG=$(docker images washing_plugin --format "{{.ID}}")
if docker tag "$IMAGE_TAG" harith2001/coconut-peat-supply-chain_core_system-washing:latest && docker push harith2001/coconut-peat-supply-chain_core_system-washing:latest; then
  echo "Docker image push as latest successful."
else
  echo "Failed to push Docker image as latest. Exiting..."
  exit 1
fi

# Apply the kube yaml file
echo "Applying kube yaml file..."
if kubectl apply -f "$KUBEFILE"; then
  echo "Kube yaml file applied successfully."
else
  echo "Failed to apply kube yaml file. Exiting..."
  exit 1
fi
