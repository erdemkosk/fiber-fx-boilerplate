# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy files
COPY . .

# Generate swagger docs
RUN swag init -g main.go

# Build app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env

EXPOSE 8090
CMD ["./main"] 