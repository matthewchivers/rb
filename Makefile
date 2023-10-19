# Set the binary name
BINARY_NAME=rb

# Build the binary
build:
	@echo "Building..."
	@go build -o $(BINARY_NAME)

# Run main.go (no build)
run:
	@echo "Running..."
	@go run main.go

# Clean up
clean:
	@echo "Cleaning..."
	@rm $(BINARY_NAME) || (echo "Failed to delete $(BINARY_NAME)"; exit 1)
	@echo "$(BINARY_NAME) deleted successfully"

# Test the project
test:
	@echo "Testing..."
	@go test ./...

# Format the code
fmt:
	@echo "Formatting..."
	@go fmt ./...

# Vet the code
vet:
	@echo "Vetting..."
	@go vet ./...

# Lint the code
lint:
	@echo "Linting..."
	@OUTPUT=$$(golint ./...); \
	if [ -z "$$OUTPUT" ]; then \
		echo "No linting errors found"; \
	else \
		echo "$$OUTPUT"; \
		exit 1; \
	fi