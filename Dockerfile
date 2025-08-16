# Multi-stage Dockerfile for OSRS OTK
# Stage 1: Build backend
FROM golang:1.24-alpine AS backend-builder

# Install dependencies for building
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy backend go mod files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy backend source code
COPY backend/cmd/ cmd/
COPY backend/internal/ internal/
COPY backend/pkg/ pkg/

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server

# Stage 2: Build frontend
FROM node:18-alpine AS frontend-builder

WORKDIR /app/frontend

# Copy package files
COPY frontend/package*.json ./

# Install dependencies (use install instead of ci for flexibility)
RUN npm install

# Copy frontend source
COPY frontend/ ./

# Build frontend
RUN npm run build

# Stage 3: Production image
FROM alpine:3.18 AS production

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy built backend binary
COPY --from=backend-builder /app/bin/server /app/server

# Copy built frontend
COPY --from=frontend-builder /app/frontend/.svelte-kit/output /app/web/frontend/build

# Copy configuration and assets
COPY backend/internal/config /app/internal/config
COPY backend/assets/ /app/assets/
COPY shared/ /app/shared/

# Create directories and set permissions
RUN mkdir -p /app/logs && \
    chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose ports
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Set environment
ENV APP_ENV=production

# Start the application
CMD ["./server"]

# Development stage (for development builds)
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
COPY shared/ ./shared/

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
COPY backend/scripts/docker-start-dev.sh /app/scripts/docker-start-dev.sh
RUN chmod +x /app/scripts/docker-start-dev.sh

# Expose ports
EXPOSE 8080 5173

# Ensure air is in PATH
ENV PATH="/go/bin:${PATH}"

# Start development with both backend and frontend
CMD ["/app/scripts/docker-start-dev.sh"]