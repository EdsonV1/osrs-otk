.PHONY: build clean dev test backend frontend docker-build docker-run docker-dev docker-prod docker-clean

# Variables
APP_NAME := osrs-otk
SERVER_NAME := $(APP_NAME)-server
FRONTEND_DIR := web/frontend
DOCKER_IMAGE := $(APP_NAME):latest
DOCKER_DEV_IMAGE := $(APP_NAME):dev

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

# Run tests with coverage
test-coverage:
	@echo "Running backend tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run integration tests
test-integration:
	@echo "Running integration tests..."
	go test -v -tags=integration ./internal/integration/...

# Run all tests (unit + integration + frontend)
test-all: test test-integration
	@echo "All tests completed"

# Run benchmarks
benchmark:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# Test with race detection
test-race:
	@echo "Running tests with race detection..."
	go test -v -race ./...

# Docker build (production)
docker-build:
	@echo "Building production Docker image..."
	docker build --target production -t $(DOCKER_IMAGE) .

# Docker build (development)
docker-build-dev:
	@echo "Building development Docker image..."
	docker build --target development -t $(DOCKER_DEV_IMAGE) .

# Docker run (single container)
docker-run:
	@echo "Running Docker container..."
	docker run --rm -p 8080:8080 -e APP_ENV=docker $(DOCKER_IMAGE)

# Docker development environment
docker-dev:
	@echo "Starting Docker development environment..."
	docker-compose -f docker-compose.dev.yml up --build

# Docker production environment
docker-prod:
	@echo "Starting Docker production environment..."
	docker-compose -f docker-compose.prod.yml up -d --build

# Docker development (detached)
docker-dev-detached:
	@echo "Starting Docker development environment (detached)..."
	docker-compose -f docker-compose.dev.yml up -d --build

# Stop Docker development
docker-dev-down:
	@echo "Stopping Docker development environment..."
	docker-compose -f docker-compose.dev.yml down

# Stop Docker production
docker-prod-down:
	@echo "Stopping Docker production environment..."
	docker-compose -f docker-compose.prod.yml down

# Docker logs
docker-logs:
	docker-compose -f docker-compose.dev.yml logs -f

# Docker clean
docker-clean:
	@echo "Cleaning Docker resources..."
	docker-compose -f docker-compose.dev.yml down -v --remove-orphans
	docker-compose -f docker-compose.prod.yml down -v --remove-orphans
	docker system prune -f
	docker volume prune -f

# Docker shell (development)
docker-shell:
	docker-compose -f docker-compose.dev.yml exec backend sh

# Docker rebuild
docker-rebuild: docker-clean docker-dev

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
	@echo ""
	@echo "Local Development:"
	@echo "  build       - Build backend and frontend"
	@echo "  backend     - Build backend only"
	@echo "  frontend    - Build frontend only"
	@echo "  deps        - Install dependencies"
	@echo "  dev         - Start development servers"
	@echo "  clean       - Clean build artifacts"
	@echo "  test        - Run tests"
	@echo ""
	@echo "Docker Development:"
	@echo "  docker-dev  - Start Docker development environment"
	@echo "  docker-dev-detached - Start Docker development (background)"
	@echo "  docker-dev-down - Stop Docker development"
	@echo "  docker-logs - Show Docker development logs"
	@echo "  docker-shell - Access development container shell"
	@echo ""
	@echo "Docker Production:"
	@echo "  docker-prod - Start Docker production environment"
	@echo "  docker-prod-down - Stop Docker production"
	@echo "  docker-build - Build production Docker image"
	@echo "  docker-run  - Run single production container"
	@echo ""
	@echo "Docker Utilities:"
	@echo "  docker-clean - Clean Docker resources"
	@echo "  docker-rebuild - Rebuild development environment"
	@echo ""
	@echo "Code Quality:"
	@echo "  fmt         - Format code"
	@echo "  lint        - Lint code"
	@echo "  help        - Show this help"