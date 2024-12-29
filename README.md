# Fiber-FX-Boilerplate

A modern Go web application boilerplate using Fiber framework with Uber FX dependency injection, MongoDB integration, and Swagger documentation.

## 🚀 Features

- **Fiber Framework**: High-performance web framework
- **Uber FX**: Dependency injection for clean architecture
- **MongoDB**: NoSQL database integration
- **Swagger**: Automated API documentation
- **Zap Logger**: Structured logging
- **Docker Support**: Containerization ready
- **GitHub Actions**: Automated swagger generation
- **Graceful Shutdown**: Proper application shutdown handling
- **Environment Config**: Easy configuration management
- **Health Check**: Built-in health monitoring

## 📁 Project Structure

```
├── .github/
│   └── workflows/       # GitHub Actions workflows
├── src/
│   ├── bootstrap/
│   │   ├── config/     # Application configuration
│   │   ├── logger/     # Zap logger setup
│   │   ├── mongodb/    # MongoDB connection
│   │   └── server/     # Fiber server setup
│   ├── internal/
│   │   ├── handler/    # HTTP request handlers
│   │   ├── middleware/ # Custom middlewares
│   │   ├── model/     # Data models
│   │   ├── repository/ # Data access layer
│   │   ├── routes/    # Route definitions
│   │   └── service/   # Business logic
├── docs/              # Swagger documentation
├── Dockerfile        # Docker configuration
├── docker-compose.yml
├── .env.example      # Environment variables template
└── main.go          # Application entry point
```

## 🛠 Prerequisites

- Go 1.21+
- MongoDB
- Docker (optional)

## ⚡️ Quick Start

1. **Clone the repository**
```bash
git clone https://github.com/yourusername/fiber-fx-boilerplate.git
```

2. **Set up environment variables**
```bash
cp .env.example .env
# Edit .env with your configurations
```

3. **Run with Docker**
```bash
docker-compose up --build
```

**OR**

3. **Run locally**
```bash
go mod download
go run main.go
```

## 🔄 API Endpoints

### Core Endpoints
- `GET /api/health` - Health check endpoint

### Foo Resource
- `GET /api/foo` - Get all items
- `GET /api/foo/:id` - Get item by ID
- `POST /api/foo` - Create new item
- `PUT /api/foo/:id` - Update existing item

## 📚 API Documentation

Swagger UI is available at:
```
http://localhost:8090/swagger/
```

Documentation is automatically generated and updated via GitHub Actions.

## 🐳 Docker Support

Build and run with Docker:
```bash
# Build image
docker build -t fiber-fx-app .

# Run container
docker run -p 8090:8090 fiber-fx-app
```

Or use docker-compose:
```bash
docker-compose up
```

## 🔧 Configuration

Configuration is managed through environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| APP_ENV | Application environment | development |
| PORT | Server port | 8090 |
| MONGO_DB_URL | MongoDB connection URL | - |
| LOG_LEVEL | Logging level | info |
| QUERY_TIMEOUT | Database query timeout | 1m |

## 🏗 Architecture

The project follows clean architecture principles:

- **Handlers**: HTTP request handling
- **Services**: Business logic implementation
- **Repositories**: Data access abstraction
- **Models**: Data structures
- **Middleware**: Request/Response processing

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [Uber FX](https://github.com/uber-go/fx)
- [Swaggo](https://github.com/swaggo/swag)


