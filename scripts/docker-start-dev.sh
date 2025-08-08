#!/bin/sh

# Start script for development Docker container
echo "Starting OSRS OTK development environment..."

# Start backend with air (hot reload) in background
cd /app
air -c .air.toml &
BACKEND_PID=$!

# Start frontend dev server
cd /app/web/frontend
npm run dev -- --host 0.0.0.0 &
FRONTEND_PID=$!

# Function to cleanup on exit
cleanup() {
    echo "Shutting down services..."
    kill $BACKEND_PID 2>/dev/null || true
    kill $FRONTEND_PID 2>/dev/null || true
    exit 0
}

# Trap cleanup function on script exit
trap cleanup EXIT INT TERM

echo "✅ Development environment started!"
echo "Backend API: http://localhost:8080"
echo "Frontend: http://localhost:5173"

# Wait for any process to exit
wait