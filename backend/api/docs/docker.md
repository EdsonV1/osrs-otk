# Docker Deployment Guide

This guide covers how to run OSRS OTK using Docker for both development and production environments.

## 🚀 Quick Start

### Development Environment
```bash
# Start development environment with hot reload
make docker-dev

# Or use the helper script
./scripts/docker-dev.sh
```

### Production Environment
```bash
# Start production environment
make docker-prod

# Or single container
make docker-build
make docker-run
```

## 📋 Prerequisites

- Docker 20.10+
- Docker Compose 2.0+

## 🏗️ Architecture

### Multi-Stage Dockerfile

The `Dockerfile` uses a multi-stage build approach:

1. **Backend Builder**: Builds the Go application
2. **Frontend Builder**: Builds the SvelteKit frontend
3. **Production**: Minimal Alpine Linux image with both applications
4. **Development**: Development image with hot reload capabilities

### Docker Compose Environments

#### Development (`docker-compose.dev.yml`)
- Backend with hot reload using Air
- Frontend with Vite dev server
- Volume mounts for live code changes
- Separate containers for backend and frontend

#### Production (`docker-compose.prod.yml`)
- Single optimized container
- Nginx reverse proxy
- Health checks
- Logging configuration
- Production-optimized settings

## 🔧 Available Commands

### Development Commands
```bash
make docker-dev              # Start development environment
make docker-dev-detached     # Start in background
make docker-dev-down         # Stop development environment
make docker-logs             # View logs
make docker-shell            # Access container shell
```

### Production Commands
```bash
make docker-prod             # Start production environment
make docker-prod-down        # Stop production environment
make docker-build            # Build production image
make docker-run              # Run single container
```

### Utility Commands
```bash
make docker-clean            # Clean all Docker resources
make docker-rebuild          # Rebuild development environment
```

## 🌐 Service Ports

| Service | Development | Production |
|---------|-------------|------------|
| Frontend | 5173 | - |
| Backend API | 8080 | 8080 |
| Nginx (Prod) | - | 80, 443 |

## 🔒 Environment Configuration

### Environment Variables

The application reads configuration from:
1. Environment variables
2. YAML configuration files
3. Docker-specific overrides

Key environment variables:
- `APP_ENV`: Environment (development/production/docker)
- `SERVER_PORT`: Server port (default: 8080)
- `SERVER_HOST`: Server host (default: 0.0.0.0 in Docker)

### Configuration Files

Docker uses specific configuration files:
- `internal/config/environments/docker.yaml`: Docker-specific settings
- `internal/config/environments/production.yaml`: Production settings

## 🚀 Production Deployment

### Single Container Deployment
```bash
# Build production image
docker build -t osrs-otk:latest .

# Run container
docker run -d \
  --name osrs-otk \
  -p 8080:8080 \
  -e APP_ENV=production \
  --restart unless-stopped \
  osrs-otk:latest
```

### Docker Compose Deployment
```bash
# Start production stack with Nginx
docker-compose -f docker-compose.prod.yml up -d

# Check status
docker-compose -f docker-compose.prod.yml ps

# View logs
docker-compose -f docker-compose.prod.yml logs -f
```

### Health Checks

The production image includes health checks:
```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/skill-data/hunter || exit 1
```

## 🔄 Development Workflow

### Starting Development
```bash
# Clone repository
git clone <repository-url>
cd osrs-otk

# Start Docker development environment
make docker-dev
```

### Making Changes
- Backend changes trigger automatic rebuild (via Air)
- Frontend changes trigger automatic rebuild (via Vite HMR)
- Configuration changes require container restart

### Debugging
```bash
# View logs
make docker-logs

# Access backend container
make docker-shell

# Restart specific service
docker-compose -f docker-compose.dev.yml restart backend
```

## 📊 Monitoring & Logging

### Production Logging
```bash
# View application logs
docker-compose -f docker-compose.prod.yml logs app

# View Nginx logs
docker-compose -f docker-compose.prod.yml logs nginx

# Follow logs in real-time
docker-compose -f docker-compose.prod.yml logs -f
```

### Log Rotation
Production containers use JSON file logging with rotation:
```yaml
logging:
  driver: "json-file"
  options:
    max-size: "10m"
    max-file: "3"
```

## 🛠️ Troubleshooting

### Common Issues

#### Container Won't Start
```bash
# Check container logs
docker logs osrs-otk

# Check if ports are in use
lsof -i :8080
```

#### Database/File Permissions
```bash
# Check file permissions
docker exec osrs-otk ls -la /app/assets

# Fix permissions if needed
docker exec osrs-otk chown -R appuser:appgroup /app
```

#### Network Issues
```bash
# Check network connectivity
docker exec osrs-otk ping google.com

# Check internal service connectivity
docker exec osrs-otk wget -O- http://localhost:8080/api/skill-data/hunter
```

### Performance Tuning

#### Memory Limits
```yaml
services:
  app:
    mem_limit: 512m
    memswap_limit: 512m
```

#### CPU Limits
```yaml
services:
  app:
    cpus: '0.5'
```

## 🔐 Security Considerations

### Production Security
- Non-root user in container
- Minimal base image (Alpine Linux)
- Security headers via Nginx
- Rate limiting configured
- No secrets in image layers

### Network Security
```yaml
# Restrict external access
networks:
  internal:
    driver: bridge
    internal: true
```

## 📈 Scaling

### Horizontal Scaling
```yaml
services:
  app:
    deploy:
      replicas: 3
    
  nginx:
    depends_on:
      - app
```

### Load Balancing
Configure Nginx upstream for multiple backend instances:
```nginx
upstream backend {
    server app_1:8080;
    server app_2:8080;
    server app_3:8080;
}
```

## 🚢 CI/CD Integration

### GitHub Actions Example
```yaml
- name: Build Docker Image
  run: docker build -t osrs-otk:${{ github.sha }} .

- name: Run Tests in Docker
  run: |
    docker run --rm osrs-otk:${{ github.sha }} go test ./...
```

### Deployment Pipeline
1. Build multi-stage image
2. Run tests in container
3. Push to registry
4. Deploy to production
5. Health check validation

This Docker setup provides a complete containerization solution for OSRS OTK, supporting both development and production workflows with industry best practices.