# Go parameters
BINARY_NAME=fiber-app
MAIN_FILE=server/main.go

# Build variables
BUILD_DIR=build
BINARY_WINDOWS=$(BINARY_NAME).exe
BINARY_LINUX=$(BINARY_NAME)

.PHONY: all build run clean test coverage dev lint vet

all: test build

build:
	@echo "Building..."
	go build -o $(BUILD_DIR)/$(BINARY_WINDOWS) $(MAIN_FILE)

run:
	@echo "Running..."
	go run $(MAIN_FILE)

clean:
	@echo "Cleaning..."
	go clean
	rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	go test -v ./...

coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

dev:
	@echo "Running in development mode..."
	air

lint:
	@echo "Running linter..."
	golangci-lint run

vet:
	@echo "Running go vet..."
	go vet ./...

build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_LINUX) $(MAIN_FILE)

build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_WINDOWS) $(MAIN_FILE)

install-tools:
	@echo "Installing development tools..."
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
