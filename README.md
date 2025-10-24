# Coubcore Blockchain

A complete blockchain implementation with Go backend and React frontend.

## Overview

Coubcore is a fully functional blockchain system that includes:

1. **Go Backend**: A robust blockchain implementation with Proof of Work consensus
2. **React Frontend**: A modern web interface for wallet management and blockchain exploration
3. **Smart Contracts**: A simple virtual machine for executing smart contracts
4. **Deployment**: Docker, Kubernetes, and CI/CD configurations for production deployment
5. **Monitoring**: Prometheus and Grafana integration for system monitoring

## Features

### Blockchain Core
- Block structure with SHA256 hashing
- Proof of Work consensus mechanism
- UTXO-based transaction system
- Peer-to-peer networking
- JSON-RPC API

### Wallet System
- Cryptographic wallet creation
- Balance management
- Transaction sending and receiving
- QR code generation

### Smart Contracts
- Simple virtual machine
- Gas calculation system
- Example contracts

### Frontend
- Modern React/TypeScript interface
- Wallet management
- Blockchain explorer
- Responsive design
- Dark theme

### Deployment
- Docker containerization
- Kubernetes manifests
- CI/CD pipeline with GitHub Actions
- Makefile for common tasks

### Monitoring
- Prometheus metrics
- Health and readiness checks
- Structured logging

## Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                            Frontend                                 │
│  ┌─────────────┐  ┌──────────────────┐  ┌────────────────────────┐ │
│  │   Wallet    │  │ Blockchain       │  │ Advanced Features      │ │
│  │             │  │ Explorer         │  │                        │ │
│  │ - Creation  │  │                  │  │ - Theme Support        │ │
│  │ - Balance   │  │ - Block List     │  │ - Language Support     │ │
│  │ - Send TX   │  │ - Block Detail   │  │ - Export Functionality │ │
│  └─────────────┘  │ - Search         │  │ - Hardware Wallet      │ │
│                   │ - Statistics     │  │   Integration          │ │
│                   └──────────────────┘  └────────────────────────┘ │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                            Backend                                  │
│  ┌─────────────┐  ┌──────────────────┐  ┌────────────────────────┐ │
│  │   Core      │  │ Consensus &      │  │ Networking & API       │ │
│  │             │  │ Cryptography     │  │                        │ │
│  │ - Blocks    │  │                  │  │ - P2P Networking       │ │
│  │ - Chain     │  │ - Proof of Work  │  │ - JSON-RPC API         │ │
│  │ - Storage   │  │ - Wallet System  │  │ - WebSocket Support    │ │
│  └─────────────┘  │ - Transactions   │  │ - Health Checks        │ │
│                   └──────────────────┘  └────────────────────────┘ │
│                                    │                              │
│                                    ▼                              │
│                   ┌─────────────────────────────────┐             │
│                   │        Smart Contracts          │             │
│                   │                                 │             │
│                   │ - Virtual Machine               │             │
│                   │ - Gas System                    │             │
│                   │ - Example Contracts             │             │
│                   └─────────────────────────────────┘             │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                           Deployment                                │
│  ┌─────────────┐  ┌──────────────────┐  ┌────────────────────────┐ │
│  │   Docker    │  │ Kubernetes       │  │ CI/CD                  │ │
│  │             │  │                  │  │                        │ │
│  │ - Backend   │  │ - Manifests      │  │ - GitHub Actions       │ │
│  │ - Frontend  │  │ - Services       │  │ - Automated Testing    │ │
│  │ - Compose   │  │ - Ingress        │  │ - Deployment Pipeline  │ │
│  └─────────────┘  └──────────────────┘  └────────────────────────┘ │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                           Monitoring                                │
│  ┌─────────────┐  ┌──────────────────┐  ┌────────────────────────┐ │
│  │ Prometheus  │  │ Grafana          │  │ Logging                │ │
│  │             │  │                  │  │                        │ │
│  │ - Metrics   │  │ - Dashboards     │  │ - Structured Logging   │ │
│  │ - Alerts    │  │ - Visualization  │  │ - Log Aggregation      │ │
│  │ - Exporters │  │ - Monitoring     │  │ - Error Tracking       │ │
│  └─────────────┘  └──────────────────┘  └────────────────────────┘ │
└─────────────────────────────────────────────────────────────────────┘
```

## Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- Docker (for containerization)
- Kubernetes (for deployment)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/antontuzov/coubcore.git
   cd coubcore
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Install frontend dependencies:
   ```bash
   cd frontend
   npm install
   cd ..
   ```

### Running the Application

#### Development Mode

1. Start the Go backend:
   ```bash
   go run cmd/coubcore/main.go
   ```

2. Start the React frontend:
   ```bash
   cd frontend
   npm start
   ```

#### Production Mode

1. Build and run with Docker:
   ```bash
   docker-compose up
   ```

2. Or deploy to Kubernetes:
   ```bash
   kubectl apply -f k8s/
   ```

### Using Makefile

The project includes a Makefile for common tasks:

```bash
# Build the Go backend
make build

# Run the Go backend
make run

# Run tests
make test

# Run frontend development server
make frontend-dev

# Build Docker images
make docker-build

# Deploy to Kubernetes
make k8s-deploy

# Show all available commands
make help
```

## API Endpoints

### Blockchain Info
- `GET /api/v1/info` - Get blockchain information

### Blocks
- `GET /api/v1/block?index={index}` - Get block by index

### Transactions
- `GET /api/v1/transaction?txid={txid}` - Get transaction by ID
- `GET /api/v1/balance?address={address}` - Get balance for address
- `POST /api/v1/send` - Send a transaction

### Network
- `GET /api/v1/peers` - Get connected peers

### Health Checks
- `GET /health` - Health check endpoint
- `GET /ready` - Readiness check endpoint

### Metrics
- `GET /metrics` - Prometheus metrics endpoint

## Project Structure

```
coubcore/
├── cmd/
│   └── coubcore/          # Main application entry point
├── internal/
│   ├── blockchain/        # Blockchain core implementation
│   │   ├── api/           # JSON-RPC API
│   │   ├── consensus/      # Consensus mechanisms
│   │   ├── core/           # Core blockchain components
│   │   ├── network/        # P2P networking
│   │   └── wallet/         # Wallet system
│   ├── logging/            # Logging utilities
│   ├── metrics/            # Prometheus metrics
│   └── vm/                # Virtual machine for smart contracts
├── frontend/              # React frontend
│   ├── public/            # Static assets
│   ├── src/               # Source code
│   │   ├── components/     # React components
│   │   ├── contexts/       # React contexts
│   │   ├── services/       # API services
│   │   ├── types/          # TypeScript types
│   │   └── utils/          # Utility functions
│   ├── Dockerfile          # Frontend Dockerfile
│   └── nginx.conf          # Nginx configuration
├── k8s/                   # Kubernetes manifests
├── prometheus/            # Prometheus configuration
├── tests/                 # Test files
├── .github/workflows/     # CI/CD pipelines
├── Dockerfile             # Backend Dockerfile
├── docker-compose.yml     # Docker Compose configuration
├── Makefile               # Build automation
├── go.mod                 # Go module dependencies
└── README.md              # This file
```

## Testing

### Go Tests
```bash
go test ./...
```

### Frontend Tests
```bash
cd frontend
npm test
```

## Deployment

### Docker
```bash
# Build images
docker build -t coubcore-backend .
docker build -t coubcore-frontend ./frontend

# Run with Docker Compose
docker-compose up
```

### Kubernetes
```bash
# Deploy to Kubernetes
kubectl apply -f k8s/

# Check deployment status
kubectl get pods
```

## Monitoring

### Prometheus
Prometheus is configured to scrape metrics from the backend service.

### Grafana
Grafana dashboards provide visualization of blockchain metrics.

### Health Checks
Health and readiness endpoints provide status information for monitoring.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Go](https://golang.org/)
- [React](https://reactjs.org/)
- [Prometheus](https://prometheus.io/)
- [Grafana](https://grafana.com/)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)