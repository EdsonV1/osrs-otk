.PHONY: build clean dev test backend frontend docker-build docker-run

# Variables
APP_NAME := osrs-otk
SERVER_NAME := $(APP_NAME)-server
FRONTEND_DIR := web/frontend
DOCKER_IMAGE := $(APP_NAME):latest

# Default target
all: build

# Build everything
build: backend frontend

# Build backend
backend:
	@echo "Building backend..."
	go build -o bin/$(SERVER_NAME) ./cmd/server

# Build frontend
frontend:
	@echo "Building frontend..."
	cd $(FRONTEND_DIR) && npm run build

# Install dependencies
deps:
	@echo "Installing backend dependencies..."
	go mod tidy
	@echo "Installing frontend dependencies..."
	cd $(FRONTEND_DIR) && npm install

# Development mode
dev:
	@echo "Starting development servers..."
	@make backend &
	@cd $(FRONTEND_DIR) && npm run dev &
	@APP_ENV=development ./bin/$(SERVER_NAME) &
	@echo "Development servers started"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -rf $(FRONTEND_DIR)/build/
	rm -rf $(FRONTEND_DIR)/.svelte-kit/

# Run tests
test:
	@echo "Running backend tests..."
	go test -v ./...
	@echo "Running frontend tests..."
	cd $(FRONTEND_DIR) && npm test

# Docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Docker run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 -p 5173:5173 $(DOCKER_IMAGE)

# Format code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

# Lint code  
lint:
	@echo "Linting Go code..."
	golangci-lint run
	@echo "Linting frontend code..."
	cd $(FRONTEND_DIR) && npm run lint

# Help
help:
	@echo "Available targets:"
	@echo "  build       - Build backend and frontend"
	@echo "  backend     - Build backend only"
	@echo "  frontend    - Build frontend only"
	@echo "  deps        - Install dependencies"
	@echo "  dev         - Start development servers"
	@echo "  clean       - Clean build artifacts"
	@echo "  test        - Run tests"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run  - Run Docker container"
	@echo "  fmt         - Format code"
	@echo "  lint        - Lint code"
	@echo "  help        - Show this help"