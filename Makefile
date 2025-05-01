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


# Docker image settings
DOCKER_IMAGE_NAME := mailsend-server
DOCKERFILE_PATH := cmd/server/Dockerfile
DOCKER_CONTEXT := .

.PHONY: docker-build docker-run docker-clean

# Build Docker image
docker-build:
	@echo "🐳 Building Docker image..."
	@docker build -f $(DOCKERFILE_PATH) -t $(DOCKER_IMAGE_NAME) $(DOCKER_CONTEXT)

# Run Docker container
docker-run:
	@echo "🚀 Running Docker container..."
	@docker run --rm -it -p 8080:8080 \
		-v $(PWD)/runtime:/root/runtime \
		$(DOCKER_IMAGE_NAME)

# Remove Docker image
docker-clean:
	@echo "🧹 Removing Docker image..."
	@docker rmi $(DOCKER_IMAGE_NAME) || true
