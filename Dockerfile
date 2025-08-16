# Development-only Dockerfile for OSRS OTK
FROM golang:1.24-alpine AS development

# Install system dependencies
RUN apk add --no-cache git ca-certificates tzdata curl nodejs npm

WORKDIR /app

# Set Go environment for module downloads
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

# Install air for hot reloading with explicit version
RUN go install github.com/cosmtrek/air@v1.49.0

# Copy source code first
COPY backend/ ./backend/
COPY frontend/ ./frontend/
COPY .air.toml ./

# Download Go dependencies from backend directory
WORKDIR /app/backend
RUN go mod download
WORKDIR /app

# Install frontend dependencies
WORKDIR /app/frontend
RUN npm install

WORKDIR /app

# Create scripts directory and copy the development start script
RUN mkdir -p /app/scripts
COPY scripts/docker-start-dev.sh /app/scripts/docker-start-dev.sh
RUN chmod +x /app/scripts/docker-start-dev.sh

# Expose ports
EXPOSE 8080 5173

# Ensure air is in PATH
ENV PATH="/go/bin:${PATH}"

# Start development with both backend and frontend
CMD ["/app/scripts/docker-start-dev.sh"]