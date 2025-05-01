# Project settings
BIN_DIR := bin
CLIENT_BIN := $(BIN_DIR)/client
SERVER_BIN := $(BIN_DIR)/server
EXAMPLE_BIN := $(BIN_DIR)/example
CERT_CONF := scripts/cert.conf
CERT_GEN := scripts/generate_cert.sh

# Default target
.PHONY: all
all: build

# Build both client and server
.PHONY: build
build: $(CLIENT_BIN) $(SERVER_BIN) $(EXAMPLE_BIN)

# Build client binary
$(CLIENT_BIN): cmd/client/*.go pkg/**/**/*.go
	@echo "🔨 Building client..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(CLIENT_BIN) ./cmd/client

# Build server binary
$(SERVER_BIN): cmd/server/*.go pkg/**/**/*.go
	@echo "🔨 Building server..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(SERVER_BIN) ./cmd/server

# Build example binary
$(EXAMPLE_BIN): cmd/example/*.go pkg/**/**/*.go
	@echo "🔨 Building example..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(EXAMPLE_BIN) ./cmd/example

# Generate TLS certs
.PHONY: generate-cert
generate-cert: $(CERT_CONF) $(CERT_GEN)
	@echo "🔐 Generating TLS certificate and key..."
	@bash $(CERT_GEN)

# Clean binaries and certs
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(BIN_DIR) 


# Docker Compose targets
.PHONY: docker-build docker-up docker-down docker-clean

docker-build:
	@echo "🐳 Building Docker image with Compose..."
	@docker compose build

docker-up:
	@echo "🚀 Starting services with Compose..."
	@docker compose up

docker-down:
	@echo "🛑 Stopping services with Compose..."
	@docker compose down

docker-clean: docker-down
	@echo "🧹 Removing built Docker image..."
	@docker rmi mailsend-server || true
