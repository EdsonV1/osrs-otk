#!/bin/bash

# Docker development environment startup script
set -e

echo "🐳 Starting OSRS OTK Docker development environment..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check prerequisites
echo -e "${YELLOW}Checking prerequisites...${NC}"

if ! command_exists docker; then
    echo -e "${RED}❌ Docker not found. Please install Docker.${NC}"
    exit 1
fi

if ! command_exists docker-compose; then
    echo -e "${RED}❌ Docker Compose not found. Please install Docker Compose.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Prerequisites check passed${NC}"

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}❌ Docker daemon is not running. Please start Docker.${NC}"
    exit 1
fi

# Build and start services
echo -e "${BLUE}🔨 Building and starting Docker services...${NC}"
docker-compose -f docker-compose.dev.yml up --build

# Cleanup function
cleanup() {
    echo -e "\n${YELLOW}🛑 Shutting down Docker services...${NC}"
    docker-compose -f docker-compose.dev.yml down
    exit 0
}

# Trap cleanup function on script exit
trap cleanup EXIT INT TERM

echo -e "${GREEN}✅ Docker development environment ready!${NC}"
echo ""
echo -e "${GREEN}🌐 Frontend: http://localhost:5173${NC}"
echo -e "${GREEN}🔧 Backend API: http://localhost:8080${NC}"
echo ""
echo -e "${YELLOW}Press Ctrl+C to stop all services${NC}"

# Wait for services
wait