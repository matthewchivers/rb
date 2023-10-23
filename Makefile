# Set the binary name
BINARY_NAME=rb

# Build the binary to the bin directory
build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME)

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
	@OUTPUT=$$(go fmt ./...); \
	if [ -z "$$OUTPUT" ]; then \
		echo "No formatting errors found"; \
	else \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

# Vet the code
vet:
	@echo "Vetting..."
	@OUTPUT=$$(go vet ./...); \
	if [ -z "$$OUTPUT" ]; then \
		echo "No vetting errors found"; \
	else \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

# Lint the code
lint:
	@echo "Linting..."
	@command -v golint >/dev/null 2>&1 || { echo >&2 "golint is required but not installed. Aborting. Install with go get -u golang.org/x/lint/golint"; exit 1; }
	@OUTPUT=$$(golint ./...); \
	if [ -z "$$OUTPUT" ]; then \
		echo "No linting errors found"; \
	else \
		echo "$$OUTPUT"; \
		exit 1; \
	fi