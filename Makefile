# Variables
BINARY_NAME=murmur
VERSION?=dev
COMMIT?=unknown
DATE?=unknown
BUILD_FLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

# Default target
.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development targets
.PHONY: run
run: ## Run the application in development mode
	go run ./cmd/murmur/main.go

.PHONY: run-interactive
run-interactive: ## Run in interactive mode
	go run ./cmd/murmur/main.go

.PHONY: run-with-prompt
run-with-prompt: ## Run with a specific prompt (set PROMPT="your prompt")
	@if [ -z "$(PROMPT)" ]; then \
		echo "Usage: make run-with-prompt PROMPT=\"your prompt here\""; \
		exit 1; \
	fi
	go run ./cmd/murmur/main.go -prompt="$(PROMPT)"

# Build targets
.PHONY: build
build: ## Build the binary
	go build $(BUILD_FLAGS) -o $(BINARY_NAME) ./cmd/murmur

.PHONY: build-linux
build-linux: ## Build for Linux
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-linux ./cmd/murmur

.PHONY: build-windows
build-windows: ## Build for Windows
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-windows.exe ./cmd/murmur

.PHONY: build-darwin
build-darwin: ## Build for macOS
	GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-darwin ./cmd/murmur

.PHONY: build-all
build-all: build-linux build-windows build-darwin ## Build for all platforms

# Test targets
.PHONY: test
test: ## Run tests
	go test -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: test-race
test-race: ## Run tests with race detection
	go test -race ./...

# Code quality targets
.PHONY: fmt
fmt: ## Format code
	go fmt ./...

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: lint
lint: ## Run golangci-lint (requires golangci-lint to be installed)
	golangci-lint run

.PHONY: tidy
tidy: ## Tidy go modules
	go mod tidy

.PHONY: check
check: fmt vet test ## Run all checks (format, vet, test)

# Docker targets
.PHONY: docker-build
docker-build: ## Build Docker image
	docker build -t $(BINARY_NAME):$(VERSION) .

.PHONY: docker-run
docker-run: ## Run Docker container
	docker run -it --rm \
		-e OPENAI_API_KEY=$(OPENAI_API_KEY) \
		$(BINARY_NAME):$(VERSION)

.PHONY: docker-dev
docker-dev: ## Run development Docker container
	docker-compose --profile dev up murmur-dev

# Clean targets
.PHONY: clean
clean: ## Clean build artifacts
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-*
	rm -f coverage.out coverage.html

.PHONY: clean-deps
clean-deps: ## Clean dependencies
	go clean -modcache

# Install targets
.PHONY: install
install: build ## Install binary to GOPATH/bin
	go install $(BUILD_FLAGS) ./cmd/murmur

.PHONY: uninstall
uninstall: ## Remove binary from GOPATH/bin
	go clean -i ./cmd/murmur

# Development setup
.PHONY: setup
setup: ## Setup development environment
	go mod download
	go mod verify
	@echo "Development environment setup complete!"
	@echo "Don't forget to:"
	@echo "  1. Copy .env.example to .env"
	@echo "  2. Set your OPENAI_API_KEY in .env"

# Release targets
.PHONY: release
release: clean build-all ## Create release builds
	@echo "Release builds created:"
	@ls -la $(BINARY_NAME)-*

.PHONY: version
version: ## Show version information
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT)"
	@echo "Date: $(DATE)"