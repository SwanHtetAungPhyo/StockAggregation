
.PHONY: all build run clean

# Set the Go build main file 
MAIN_FILE=main.go
BINARY_NAME=./bin/stock_aggregator

# Build the Go application
build:
	@echo "Building the application..."
	GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -o $(BINARY_NAME) $(MAIN_FILE)

# Run the built application
run: build
	@echo "Running the application..."
	./$(BINARY_NAME)

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# Default target to build and run the application
all: run