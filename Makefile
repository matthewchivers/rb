# Set the binary name
BINARY_NAME := rb

# Define Go related variables
GO := go
GO_BUILD_FLAGS := -o bin/$(BINARY_NAME)
GOLINT := golint

# Define color codes
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
MAGENTA := \033[0;35m
CYAN := \033[0;36m
RESET := \033[0m
BOLD := \033[1m

.PHONY: build run clean test fmt vet lint modtidy

# Build the binary to the bin directory
build: fmt test vet lint modtidy
	@echo "$(BOLD)$(GREEN)Building...$(RESET)"
	@$(GO) build $(GO_BUILD_FLAGS); \
	if [ $$? -eq 0 ]; then \
		echo "Build successful"; \
	fi

# Run main.go (no build)
run: fmt test vet lint modtidy
	@echo "$(BOLD)$(GREEN)Running...$(RESET)"
	@$(GO) run main.go

# Clean up
clean:
	@echo "$(BOLD)$(RED)Cleaning...$(RESET)"
	@rm $(BINARY_NAME) && echo "$(BINARY_NAME) deleted successfully" || echo "Failed to delete $(BINARY_NAME)"

# Test the project
test:
	@echo "$(BOLD)$(CYAN)Testing...$(RESET)"
	@$(GO) test ./...; \
	if [ $$? -eq 0 ]; then \
		echo "All tests passed"; \
	fi

# Format the code
fmt:
	@echo "$(BOLD)$(MAGENTA)Formatting...$(RESET)"
	@$(GO) fmt ./...; \
	if [ $$? -eq 0 ]; then \
		echo "No formatting errors found"; \
	fi

# Vet the code
vet:
	@echo "$(BOLD)$(YELLOW)Vetting...$(RESET)"
	@$(GO) vet ./...; \
	if [ $$? -eq 0 ]; then \
		echo "No vetting errors found"; \
	fi

# Lint the code
lint:
	@echo "$(BOLD)$(BLUE)Linting...$(RESET)"
	@command -v $(GOLINT) >/dev/null 2>&1 || { echo >&2 "$(GOLINT) is required but not installed. Aborting. Install with $(GO) get -u golang.org/x/lint/golint"; exit 1; }
	@$(GOLINT) ./...; \
	if [ $$? -eq 0 ]; then \
		echo "No linting errors found"; \
	fi

# Tidy go modules
modtidy:
	@echo "$(BOLD)$(CYAN)Tidying modules...$(RESET)"
	@$(GO) mod tidy; \
	if [ $$? -eq 0 ]; then \
		echo "Modules tidied successfully"; \
	fi
