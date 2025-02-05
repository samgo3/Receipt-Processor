
MAIN=cmd/main.go
BINARY_NAME=receipt-service

# Swagger parameters
SWAGGER_OUT=docs
SWAGGER_DIR=cmd,internal

.PHONY: all build test clean run swagger help


help:
	@echo "Available commands:"
	@echo "  build   - Build the application"
	@echo "  test    - Run tests"
	@echo "  clean   - Clean up build artifacts"
	@echo "  run     - Build and run the application"
	@echo "  swagger - Generate Swagger documentation"
	@echo "  help    - Show this help message"
	@echo "  all	 - Run all tasks (test, swagger, build)"

all: test swagger build 

build:
	go build -o $(BINARY_NAME) $(MAIN)

test:
	go test -v ./...

clean:
	go mod tidy
	go clean
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

swagger:
	swag init -d $(SWAGGER_DIR) -o $(SWAGGER_OUT)


