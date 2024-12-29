# Fiber-FX-Boilerplate

A modern Go web application boilerplate using Fiber framework with Uber FX dependency injection and MongoDB integration.

## Features

- **Fiber Framework**: High-performance web framework
- **Uber FX**: Dependency injection for clean architecture
- **MongoDB**: NoSQL database integration
- **Zap Logger**: Structured logging
- **Graceful Shutdown**: Proper application shutdown handling
- **Modular Structure**: Well-organized code architecture

## Project Structure

```
├── src/
│   ├── bootstrap/
│   │   ├── config/      # Application configuration
│   │   ├── logger/      # Zap logger setup
│   │   ├── mongodb/     # MongoDB connection management
│   │   └── server/      # Fiber server initialization
│   ├── internal/
│   │   ├── model/       # Data models and DTOs
│   │   ├── handler/     # HTTP request handlers
│   │   ├── service/     # Business logic layer
│   │   ├── repository/  # Data access layer
│   │   └── routes/      # Route definitions
│   └── main.go         # Application entry point
```

## Architecture

- **Handler Layer**: Handles HTTP requests and responses
- **Service Layer**: Contains business logic
- **Repository Layer**: Manages data persistence
- **Model Layer**: Defines data structures
- **Bootstrap**: Manages application initialization

## Dependencies

- Go 1.21+
- Fiber v2
- Uber FX
- MongoDB Go Driver
- Zap Logger

## Setup & Running

1. Clone the repository
2. Configure environment variables in `.env`
3. Install dependencies:
```bash
go mod tidy
```
4. Run the application:
```bash
go run main.go
```

## API Endpoints

The application follows RESTful conventions:

- `GET /api/foo` - Get all items
- `GET /api/foo/:id` - Get item by ID
- `POST /api/foo` - Create new item
- `PUT /api/foo/:id` - Update existing item

## Key Features

1. **Dependency Injection**
   - Centralized dependency management using Uber FX
   - Clean and testable code structure

2. **MongoDB Integration**
   - Configured connection pooling
   - Repository pattern implementation

3. **Routing**
   - Modular route definitions
   - Group-based route organization

4. **Error Handling**
   - Consistent error responses
   - Structured logging

5. **Configuration**
   - Environment-based configuration
   - Easy to extend and modify

## Best Practices

- Clean Architecture principles
- Dependency Injection pattern
- Repository pattern
- Interface-based design
- Modular component structure

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License.


