# Coco Peat Supply Chain Management Core
The Coco Peat Supply Chain Management Core is a modular, plugin-based backend system that manages and orchestrates the supply chain processes of coco peat manufacturing. This repository contains the core system, written in Go, with gRPC communication and Docker containerization, deployed using K3s for scalable and efficient orchestration.

# 🌟 Features
 1.  Plugin Architecture: Modular design enabling the creation and management of supply chain steps as plugins (e.g., grading, cutting, washing, drying).
 2. gRPC Communication: Efficient and secure communication between core services and plugins.
 3. Containerized Services: Plugin and core services containerized using Docker.
 4. K3s Orchestration: Lightweight Kubernetes distribution for deploying, scaling, and managing plugins.
 5. NoSQL Database: MongoDB integration for data storage.
 6. Blockchain Transparency: Ensures integrity and traceability of the supply chain workflow.

# 🛠️ Tech Stack
1. Programming Language: Go (Golang)
2. Communication: gRPC
3. Database: MongoDB
4. Containerization: Docker
5. Orchestration: K3s
6. Intergrated with Blockchain and IOT sensors

# 🚀 Getting Started
Prerequisites
Ensure you have the following installed:

Go (v1.19 or later)
Docker (latest version)
K3s (latest version)
MongoDB (NoSQL database)

# Installation

1. Clone the repository:
git clone https://github.com/harith2001/Coconut-Peat-Supply-chain_core_system.git
cd Coconut-Peat-Supply-chain_core_system

2. Install dependencies:
go mod tidy

3. Set up Docker containers:
docker-compose up -d

4. Deploy services with K3s:
kubectl apply -f deployment.yaml

5. Run the core system:
go run main.go 
OR
air

# 📖 Usage
1. Plugin Registration:
    Register new plugins using the gRPC API.
    Define plugin parameters and steps as part of the workflow.

2. Workflow Execution:
    Execute workflows with pre-configured or custom plugins.
    Monitor execution results via MongoDB or blockchain records.
    
3. Customization:
    Edit or create new plugins through the workflow customization tool (separate repository).

# 🧪 Testing
Run unit and integration tests to ensure functionality:
go test ./tests/...

# THANK YOU !!

