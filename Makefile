# Makefile for Coubcore Blockchain

# Variables
GO_CMD = go
GO_BUILD = $(GO_CMD) build
GO_TEST = $(GO_CMD) test
GO_RUN = $(GO_CMD) run
GO_CLEAN = $(GO_CMD) clean
GO_DEPS = $(GO_CMD) mod tidy

# Binary name
BINARY_NAME = coubcore
BINARY_UNIX = $(BINARY_NAME)_unix

# Directories
FRONTEND_DIR = frontend

# Docker
DOCKER_CMD = docker
DOCKER_COMPOSE = docker-compose
DOCKER_IMAGE_BACKEND = coubcore-backend
DOCKER_IMAGE_FRONTEND = coubcore-frontend

# Kubernetes
KUBECTL = kubectl
K8S_DIR = k8s

# Default target
all: build

# Build the Go backend
build:
	$(GO_BUILD) -o $(BINARY_NAME) -v

# Run the Go backend
run:
	$(GO_RUN) ./cmd/coubcore

# Run tests
test:
	$(GO_TEST) -v ./...

# Clean build artifacts
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Install dependencies
deps:
	$(GO_DEPS)
	cd $(FRONTEND_DIR) && npm install

# Run the frontend development server
frontend-dev:
	cd $(FRONTEND_DIR) && npm start

# Build the frontend for production
frontend-build:
	cd $(FRONTEND_DIR) && npm run build

# Run tests for the frontend
frontend-test:
	cd $(FRONTEND_DIR) && npm test

# Build Docker images
docker-build:
	$(DOCKER_CMD) build -t $(DOCKER_IMAGE_BACKEND) .
	$(DOCKER_CMD) build -t $(DOCKER_IMAGE_FRONTEND) ./$(FRONTEND_DIR)

# Run Docker containers
docker-run:
	$(DOCKER_COMPOSE) up

# Stop Docker containers
docker-stop:
	$(DOCKER_COMPOSE) down

# Deploy to Kubernetes
k8s-deploy:
	$(KUBECTL) apply -f $(K8S_DIR)/pvc.yaml
	$(KUBECTL) apply -f $(K8S_DIR)/deployment.yaml
	$(KUBECTL) apply -f $(K8S_DIR)/service.yaml
	$(KUBECTL) apply -f $(K8S_DIR)/ingress.yaml

# Delete Kubernetes deployment
k8s-delete:
	$(KUBECTL) delete -f $(K8S_DIR)/ingress.yaml
	$(KUBECTL) delete -f $(K8S_DIR)/service.yaml
	$(KUBECTL) delete -f $(K8S_DIR)/deployment.yaml
	$(KUBECTL) delete -f $(K8S_DIR)/pvc.yaml

# Help
help:
	@echo "Available targets:"
	@echo "  all             - Build the Go backend (default)"
	@echo "  build           - Build the Go backend"
	@echo "  run             - Run the Go backend"
	@echo "  test            - Run Go tests"
	@echo "  clean           - Clean build artifacts"
	@echo "  deps            - Install dependencies"
	@echo "  frontend-dev    - Run the frontend development server"
	@echo "  frontend-build  - Build the frontend for production"
	@echo "  frontend-test   - Run tests for the frontend"
	@echo "  docker-build    - Build Docker images"
	@echo "  docker-run      - Run Docker containers"
	@echo "  docker-stop     - Stop Docker containers"
	@echo "  k8s-deploy      - Deploy to Kubernetes"
	@echo "  k8s-delete      - Delete Kubernetes deployment"
	@echo "  help            - Show this help message"

.PHONY: all build run test clean deps frontend-dev frontend-build frontend-test docker-build docker-run docker-stop k8s-deploy k8s-delete help