# Server
BINARY_NAME=server
BUILD_PATH=./cmd/api
MAIN_FILE=$(BUILD_PATH)/main.go

# Default target
all: dev

# Build the binary
build:
        go build -o $(BINARY_NAME) $(MAIN_FILE)

# Run the server
run: build
        ./$(BINARY_NAME)

# Clean up the binary
clean:
        rm -rf $(BINARY_NAME)

# Start the server without building
dev:
        go run $(MAIN_FILE)

# Install dependencies
deps:
        go mod tidy