#!/bin/bash

# Development startup script
set -e

echo "🚀 Starting OSRS OTK development environment..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check prerequisites
echo -e "${YELLOW}Checking prerequisites...${NC}"

if ! command_exists go; then
    echo -e "${RED}❌ Go not found. Please install Go.${NC}"
    exit 1
fi

if ! command_exists node; then
    echo -e "${RED}❌ Node.js not found. Please install Node.js.${NC}"
    exit 1
fi

if ! command_exists npm; then
    echo -e "${RED}❌ npm not found. Please install npm.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Prerequisites check passed${NC}"

# Install dependencies if needed
if [ ! -f "go.sum" ]; then
    echo -e "${YELLOW}📦 Installing Go dependencies...${NC}"
    go mod tidy
fi

if [ ! -d "web/frontend/node_modules" ]; then
    echo -e "${YELLOW}📦 Installing frontend dependencies...${NC}"
    cd web/frontend && npm install && cd ../..
fi

# Build backend
echo -e "${YELLOW}🔨 Building backend...${NC}"
mkdir -p bin
go build -o bin/osrs-otk-server ./cmd/server

# Set environment
export APP_ENV=development

# Start services
echo -e "${GREEN}🎯 Starting development servers...${NC}"

# Start backend server
echo -e "${YELLOW}Starting backend server on :8080...${NC}"
./bin/osrs-otk-server &
BACKEND_PID=$!

# Start frontend dev server
echo -e "${YELLOW}Starting frontend dev server on :5173...${NC}"
cd web/frontend && npm run dev &
FRONTEND_PID=$!
cd ../..

echo -e "${GREEN}✅ Development environment ready!${NC}"
echo ""
echo -e "${GREEN}🌐 Frontend: http://localhost:5173${NC}"
echo -e "${GREEN}🔧 Backend API: http://localhost:8080${NC}"
echo ""
echo -e "${YELLOW}Press Ctrl+C to stop all servers${NC}"

# Function to cleanup on exit
cleanup() {
    echo -e "\n${YELLOW}🛑 Shutting down servers...${NC}"
    kill $BACKEND_PID 2>/dev/null || true
    kill $FRONTEND_PID 2>/dev/null || true
    exit 0
}

# Trap cleanup function on script exit
trap cleanup EXIT INT TERM

# Wait for any process to exit
wait