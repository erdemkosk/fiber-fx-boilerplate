# Fiber-FX-Boilerplate Project Documentation

A Go project template using Fiber web framework with Uber FX for dependency injection and MongoDB integration.


## Project Structure
```
├── src/
│   ├── bootstrap/
│   │   ├── config/      # Configuration management
│   │   ├── logger/      # Zap logger configuration
│   │   ├── mongodb/     # MongoDB connection and setup
│   │   └── fiber/       # Fiber web framework setup
│   ├── internal/
│   │   ├── model/       # Data models
│   │   ├── service/     # Business logic
│   │   └── repository/  # Database operations
│   └── main.go         # Application entry point
├── .env                # Environment variables
└── go.mod             # Go module file
```


## Running the Project

1. Install the required dependencies:

```bash
go mod tidy
```

2. Run the project:

```bash
go run main.go
```     

## Example

```bash
2024-12-30 00:14:45     info    📦 MongoDB started successfully.
2024-12-30 00:14:45     info    🚀 Fiber server starting on :3000
^C2024-12-30 00:14:49   info    🛑 Shutting down Fiber server
2024-12-30 00:14:49     info    🔌 MongoDB stopping...
2024-12-30 00:14:49     info    ✅ MongoDB disconnected successfully.
```

It will support graceful shutdown.


