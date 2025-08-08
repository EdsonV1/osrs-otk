# OSRS OTK (Old School RuneScape Optimization Toolkit)

A comprehensive collection of calculators and tools for Old School RuneScape training optimization.

## 🎯 Features

- **Skills Calculator**: Training method optimization for all OSRS skills
- **Technique Calculators**: 
  - Wintertodt (Firemaking XP, Phoenix pet chances, loot simulation)
  - Birdhouse Runs (Hunter XP, nest predictions, profit calculations)
  - Ardougne Knights (Thieving efficiency and profit optimization)
- **Modern Architecture**: Clean, maintainable codebase with domain-driven design
- **Configuration Management**: Environment-based settings
- **Responsive Frontend**: Built with SvelteKit

## 🚀 Quick Start

### Prerequisites

- Go 1.19+
- Node.js 18+
- npm

### Development Setup

```bash
# Clone the repository
git clone <repository-url>
cd osrs-otk

# Copy environment files
cp .env.example .env
cp web/frontend/.env.example web/frontend/.env

# Install dependencies
make deps

# Start development servers
make dev
# OR use the convenience script
./scripts/dev.sh
```

The application will be available at:
- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080

### Building for Production

```bash
# Build everything
make build

# Build individually
make backend    # Builds to bin/osrs-otk-server
make frontend   # Builds to web/frontend/build/
```

## 📁 Project Structure

```
osrs-otk/
├── cmd/server/              # Application entrypoint
├── internal/
│   ├── config/             # Configuration management
│   ├── domain/             # Core business logic
│   │   ├── skill/          # Skills domain
│   │   └── calculator/     # Calculator techniques
│   ├── handlers/           # HTTP handlers
│   ├── repository/         # Data access layer
│   └── server/             # Server setup
├── assets/
│   ├── data/skills/        # Skill training data (YAML)
│   └── images/             # Static images
├── web/frontend/           # SvelteKit frontend
├── pkg/                    # Shared/public packages
├── scripts/                # Development scripts
└── deployments/            # Deployment configurations
```

## 🔧 API Endpoints

### Skills Data
- `GET /api/skill-data/{skill}` - Get training methods for a skill

### Calculator Tools
- `POST /api/wintertodt` - Wintertodt calculator
- `POST /api/birdhouse` - Birdhouse run calculator  
- `POST /api/ardyknights` - Ardougne Knights calculator

## ⚙️ Configuration

The application uses YAML configuration files in `internal/config/environments/`:

- `development.yaml` - Local development settings
- `production.yaml` - Production settings

Environment is controlled by `APP_ENV` environment variable (defaults to `development`).

## 🛠️ Available Make Targets

```bash
make build       # Build backend and frontend
make backend     # Build backend only
make frontend    # Build frontend only
make deps        # Install dependencies
make dev         # Start development servers
make clean       # Clean build artifacts
make test        # Run tests
make fmt         # Format code
make lint        # Lint code
make help        # Show all targets
```

## 📊 Data Management

Skill training data is stored in YAML format in `assets/data/skills/`. Each file contains:

```yaml
name_canonical: "skillname"
name_display: "Skill Name"
description: "Skill description..."
training_methods:
  - id: method_id
    name: "Method Name"
    level_required: 1
    xp_rate: 10000
    # ... additional method data
```

## 🧪 Testing

```bash
# Run backend tests
go test -v ./...

# Run frontend tests
cd web/frontend && npm test
```

## 📈 Architecture

The application follows **Domain-Driven Design** principles:

- **Domain Layer**: Core business logic and models
- **Repository Layer**: Data access abstraction
- **Service Layer**: Business operations orchestration
- **Handler Layer**: HTTP request/response handling
- **Infrastructure Layer**: External dependencies (files, databases)

## 🔄 Migration from Legacy

The new structure maintains **backward compatibility** with existing calculator APIs while providing:
- Better maintainability through separation of concerns
- Environment-based configuration
- Improved error handling and logging
- Cleaner data formats (YAML vs JSON)

## 🚢 Deployment

### Using Make
```bash
make build
./bin/osrs-otk-server
```

### Environment Variables
- `APP_ENV`: Environment (development/production)
- `SERVER_PORT`: Server port (default: 8080)
- `SERVER_HOST`: Server host (default: localhost)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run `make fmt` and `make lint`
6. Submit a pull request

## 📝 License

[Add your license information here]

## 🙏 Acknowledgments

- OSRS Wiki for training method data
- Community contributors and testers