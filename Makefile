# Set the binary name
BINARY_NAME := rb

# Define Go related variables
GO := go
GO_BUILD_FLAGS := -o bin/$(BINARY_NAME)
GOLINT := golint

.PHONY: build run clean test fmt vet lint modtidy

# Build the binary to the bin directory
build: fmt test vet lint modtidy
	@echo "Building..."
	@$(GO) build $(GO_BUILD_FLAGS); \
	if [ $$? -eq 0 ]; then \
		echo "Build successful"; \
	fi

# Run main.go (no build)
run: fmt test vet lint modtidy
	@echo "Running..."
	@$(GO) run main.go

# Clean up
clean:
	@echo "Cleaning..."
	@rm $(BINARY_NAME) && echo "$(BINARY_NAME) deleted successfully" || echo "Failed to delete $(BINARY_NAME)"

# Test the project
test:
	@echo "Testing..."
	@$(GO) test ./...; \
	if [ $$? -eq 0 ]; then \
		echo "All tests passed"; \
	fi

# Format the code
fmt:
	@echo "Formatting..."
	@$(GO) fmt ./...; \
	if [ $$? -eq 0 ]; then \
		echo "No formatting errors found"; \
	fi

# Vet the code
vet:
	@echo "Vetting..."
	@$(GO) vet ./...; \
	if [ $$? -eq 0 ]; then \
		echo "No vetting errors found"; \
	fi

# Lint the code
lint:
	@echo "Linting..."
	@command -v $(GOLINT) >/dev/null 2>&1 || { echo >&2 "$(GOLINT) is required but not installed. Aborting. Install with $(GO) get -u golang.org/x/lint/golint"; exit 1; }
	@$(GOLINT) ./...; \
	if [ $$? -eq 0 ]; then \
		echo "No linting errors found"; \
	fi

# Tidy go modules
modtidy:
	@echo "Tidying modules..."
	@$(GO) mod tidy; \
	if [ $$? -eq 0 ]; then \
		echo "Modules tidied successfully"; \
	fi
