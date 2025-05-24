#!/bin/bash

# CONFIGURATION
IMAGE_NAME="harith2001/coco-core"
IMAGE_TAG="latest"
DOCKER_COMPOSE_FILE="docker-compose.yml"
K8S_DIR="kube-config" # folder where Kubernetes YAML files are stored

# Step 3: Build docker-compose services
echo "🐳 Building docker-compose services..."
docker-compose -f $DOCKER_COMPOSE_FILE build

# Step 1: Build Docker Image
echo "🔨 Building Docker image..."
docker build -t $IMAGE_NAME:$IMAGE_TAG .

# Step 2: Push Docker Image
echo "📦 Pushing Docker image to Docker Hub..."
docker push $IMAGE_NAME:$IMAGE_TAG


# Step 4: Apply Kubernetes YAMLs
echo "☸️ Applying Kubernetes manifests..."
kubectl apply -f $K8S_DIR/

echo "✅ Deployment complete."
