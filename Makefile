GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

BINARY_NAME=fiber-fx-boilerplate
MAIN_PATH=main.go


GREEN=\033[0;32m
NC=\033[0m # No Color

.PHONY: all build test clean run deps help

all: test build

build: 
	@echo "$(GREEN)Building...$(NC)"
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

test: 
	@echo "$(GREEN)Running tests...$(NC)"
	$(GOTEST) ./... -v

test-coverage: 
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	$(GOTEST) ./... -coverprofile=coverage.out
	$(GOCMD) tool cover -html=coverage.out

clean: 
	@echo "$(GREEN)Cleaning...$(NC)"
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f coverage.out

run: build 
	@echo "$(GREEN)Running...$(NC)"
	./$(BINARY_NAME)

deps: #
	@echo "$(GREEN)Checking and updating dependencies...$(NC)"
	$(GOGET) -u ./...
	$(GOMOD) tidy

fmt: 
	@echo "$(GREEN)Formatting code...$(NC)"
	$(GOCMD) fmt ./...

lint: 
	@echo "$(GREEN)Running linter...$(NC)"
	golangci-lint run

help: 
	@echo "Kullanılabilir komutlar:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%-20s$(NC) %s\n", $$1, $$2}'

.DEFAULT_GOAL := help 