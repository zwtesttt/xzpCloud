# Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_DIR=./bin

# Service binaries
VM_BINARY=$(BINARY_DIR)/vm-service
PRODUCT_BINARY=$(BINARY_DIR)/product-service
USER_BINARY=$(BINARY_DIR)/user-service
ORDER_BINARY=$(BINARY_DIR)/order-service

# Create bin directory if it doesn't exist
$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)

# Build all services
build: $(BINARY_DIR)
	$(GOBUILD) -o $(VM_BINARY) -v ./cmd/vm
	$(GOBUILD) -o $(PRODUCT_BINARY) -v ./cmd/product
	$(GOBUILD) -o $(USER_BINARY) -v ./cmd/user
	$(GOBUILD) -o $(ORDER_BINARY) -v ./cmd/order

# Build individual services
build-vm: $(BINARY_DIR)
	$(GOBUILD) -o $(VM_BINARY) -v ./cmd/vm

build-product: $(BINARY_DIR)
	$(GOBUILD) -o $(PRODUCT_BINARY) -v ./cmd/product

build-user: $(BINARY_DIR)
	$(GOBUILD) -o $(USER_BINARY) -v ./cmd/user

build-order: $(BINARY_DIR)
	$(GOBUILD) -o $(ORDER_BINARY) -v ./cmd/order

# Run services
run-vm: build-vm
	$(VM_BINARY)

run-vm-local: build-vm
	VM_CONFIG=./config/vm-local.yaml $(VM_BINARY)

run-product: build-product
	$(PRODUCT_BINARY)

run-user: build-user
	$(USER_BINARY)

run-order: build-order
	$(ORDER_BINARY)

# Run all services in background (for development)
dev:
	@echo "Starting all services..."
	@$(MAKE) run-vm &
	@$(MAKE) run-product &
	@$(MAKE) run-user &
	@$(MAKE) run-order &
	@echo "All services started in background"

# Stop all running services
stop:
	@echo "Stopping all services..."
	@pkill -f "vm-service" || true
	@pkill -f "product-service" || true
	@pkill -f "user-service" || true
	@pkill -f "order-service" || true
	@echo "All services stopped"

# Run tests
test:
	$(GOTEST) -v ./...

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)

# Install dependencies
deps:
	$(GOCMD) mod tidy
	$(GOCMD) mod download

# Setup local Kubernetes cluster with KubeVirt
setup-local-k8s:
	chmod +x scripts/setup-local-k8s-kubevirt.sh
	./scripts/setup-local-k8s-kubevirt.sh

# Show help
help:
	@echo "Available commands:"
	@echo "  build        - Build all services"
	@echo "  build-vm     - Build VM service only"
	@echo "  build-product- Build Product service only"
	@echo "  build-user   - Build User service only"
	@echo "  build-order  - Build Order service only"
	@echo "  run-vm       - Run VM service (HTTP: 8080, gRPC: 8081)"
	@echo "  run-vm-local - Run VM service with local k8s cluster"
	@echo "  run-user     - Run User service (HTTP: 8082, gRPC: 8083)"
	@echo "  run-product  - Run Product service (HTTP: 8084, gRPC: 8085)"
	@echo "  run-order    - Run Order service (HTTP: 8086, gRPC: 8087)"
	@echo "  dev          - Run all services in background"
	@echo "  stop         - Stop all running services"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Install dependencies"
	@echo "  setup-local-k8s - Setup local Kubernetes cluster with KubeVirt"
	@echo "  help         - Show this help"
	@echo ""
	@echo "Service Ports:"
	@echo "  VM Service:      HTTP :8080, gRPC :8081"
	@echo "  User Service:    HTTP :8082, gRPC :8083"
	@echo "  Product Service: HTTP :8084, gRPC :8085"
	@echo "  Order Service:   HTTP :8086, gRPC :8087"

.PHONY: build build-vm build-product build-user build-order run-vm run-vm-local run-product run-user run-order dev stop test clean deps setup-local-k8s help 