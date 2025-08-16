#!/bin/sh

# Docker development startup script for OSRS OTK
# This script starts both the backend and frontend in development mode

echo "Starting OSRS OTK Development Environment..."

# Start frontend development server in background
echo "Starting frontend development server..."
cd /app/frontend
npm run dev -- --host 0.0.0.0 --port 5173 &
FRONTEND_PID=$!

# Wait a moment for frontend to initialize
sleep 3

# Start backend with air for hot reloading
echo "Starting backend with hot reloading..."
cd /app/backend
export APP_ENV=docker
export PATH="/go/bin:$PATH"

# Use air for hot reloading if available, otherwise run directly
if command -v air >/dev/null 2>&1; then
    echo "Using air for hot reloading..."
    air -c /app/backend/.air.toml
else
    echo "Air not found, running go run directly..."
    go run ./cmd/server
fi

# If backend exits, kill frontend too
kill $FRONTEND_PID 2>/dev/null