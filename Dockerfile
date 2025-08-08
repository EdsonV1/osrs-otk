# Multi-stage Dockerfile for OSRS OTK
# Stage 1: Build backend
FROM golang:1.21-alpine AS backend-builder

# Install dependencies for building
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server

# Stage 2: Build frontend
FROM node:18-alpine AS frontend-builder

WORKDIR /app/frontend

# Copy package files
COPY web/frontend/package*.json ./

# Install dependencies
RUN npm ci

# Copy frontend source
COPY web/frontend/ ./

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
COPY --from=frontend-builder /app/frontend/build /app/web/frontend/build

# Copy configuration and assets
COPY internal/config /app/internal/config
COPY assets/ /app/assets/

# Create directories and set permissions
RUN mkdir -p /app/logs && \
    chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose ports
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/skill-data/hunter || exit 1

# Set environment
ENV APP_ENV=production

# Start the application
CMD ["./server"]

# Development stage (for development builds)
FROM golang:1.21-alpine AS development

RUN apk add --no-cache git ca-certificates tzdata curl

WORKDIR /app

# Install air for hot reloading
RUN go install github.com/cosmtrek/air@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install frontend dependencies
WORKDIR /app/web/frontend
RUN npm install

WORKDIR /app

# Expose ports
EXPOSE 8080 5173

# Start development with hot reload
CMD ["air", "-c", ".air.toml"]