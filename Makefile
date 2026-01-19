# ============================================================================
# Harness Terraform Provider - Makefile
# ============================================================================
#
# Usage: make <target>
#
# Run 'make help' for a list of available targets
#

SHELL := /bin/bash

# ============================================================================
# Configuration
# ============================================================================

# Provider details
HOSTNAME     := registry.terraform.io
NAMESPACE    := harness
NAME         := harness
BINARY       := terraform-provider-$(NAME)
VERSION      ?= 0.99.0-dev

# Build configuration
GO           := go
GOPATH       := $(shell go env GOPATH)
GOOS         := $(shell go env GOOS)
GOARCH       := $(shell go env GOARCH)
CGO_ENABLED  := 0

# Terraform plugin directory
TF_PLUGIN_DIR := ~/.terraform.d/plugins/$(HOSTNAME)/$(NAMESPACE)/$(NAME)/$(VERSION)/$(GOOS)_$(GOARCH)

# Test configuration
TEST         ?= ./...
TESTARGS     ?=
SWEEP_DIR    ?= ./internal/sweep
SWEEP_RUN    ?= all
ACC_TEST_TIMEOUT ?= 120m
UNIT_TEST_TIMEOUT ?= 10m

# Tools
TFPLUGINDOCS := $(GOPATH)/bin/tfplugindocs
CHANGELOG_BUILD := $(GOPATH)/bin/changelog-build
GOLANGCI_LINT := $(GOPATH)/bin/golangci-lint

# ============================================================================
# Colors and Output Formatting
# ============================================================================

# Colors
COLOR_RESET  := \033[0m
COLOR_BOLD   := \033[1m
COLOR_DIM    := \033[2m
COLOR_RED    := \033[0;31m
COLOR_GREEN  := \033[0;32m
COLOR_YELLOW := \033[0;33m
COLOR_BLUE   := \033[0;34m
COLOR_PURPLE := \033[0;35m
COLOR_CYAN   := \033[0;36m

# Output helpers
define log_info
	@printf "$(COLOR_BLUE)→$(COLOR_RESET) $(1)\n"
endef

define log_success
	@printf "$(COLOR_GREEN)✓$(COLOR_RESET) $(1)\n"
endef

define log_warn
	@printf "$(COLOR_YELLOW)⚠$(COLOR_RESET) $(1)\n"
endef

define log_error
	@printf "$(COLOR_RED)✗$(COLOR_RESET) $(1)\n"
endef

define log_header
	@printf "\n$(COLOR_BOLD)$(COLOR_CYAN)$(1)$(COLOR_RESET)\n"
	@printf "$(COLOR_DIM)─────────────────────────────────────$(COLOR_RESET)\n"
endef

# ============================================================================
# Default Target
# ============================================================================

.DEFAULT_GOAL := help

# ============================================================================
# Help
# ============================================================================

.PHONY: help
help: ## Show this help message
	@printf "\n$(COLOR_BOLD)$(COLOR_CYAN)  Harness Terraform Provider$(COLOR_RESET)\n"
	@printf "$(COLOR_DIM)  ─────────────────────────────────────$(COLOR_RESET)\n\n"
	@printf "  $(COLOR_BOLD)Usage:$(COLOR_RESET) make $(COLOR_GREEN)<target>$(COLOR_RESET)\n\n"
	@awk 'BEGIN {FS = ":.*##"; section=""} \
		/^##@/ { section=substr($$0, 5); next } \
		/^[a-zA-Z_-]+:.*?##/ { \
			if (section != "" && section != lastsection) { \
				printf "  $(COLOR_BOLD)%s:$(COLOR_RESET)\n", section; \
				lastsection = section \
			} \
			printf "    $(COLOR_GREEN)%-18s$(COLOR_RESET) %s\n", $$1, $$2 \
		}' $(MAKEFILE_LIST)
	@printf "\n"

# ============================================================================
# Build Targets
# ============================================================================

##@ Build

.PHONY: build
build: ## Build the provider binary
	$(call log_header,Building Provider)
	$(call log_info,Building $(BINARY) for $(GOOS)/$(GOARCH)...)
	@CGO_ENABLED=$(CGO_ENABLED) $(GO) build -trimpath -ldflags="-s -w" -o $(BINARY)
	$(call log_success,Built $(BINARY) successfully)

.PHONY: build-all
build-all: ## Build for all supported platforms
	$(call log_header,Building for All Platforms)
	@for os in darwin linux windows; do \
		for arch in amd64 arm64; do \
			printf "$(COLOR_BLUE)→$(COLOR_RESET) Building for $$os/$$arch...\n"; \
			GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 $(GO) build -trimpath -ldflags="-s -w" -o dist/$(BINARY)_$${os}_$${arch} || exit 1; \
		done; \
	done
	$(call log_success,All builds completed in dist/)

.PHONY: install
install: build ## Install provider to local Terraform plugins directory
	$(call log_header,Installing Provider Locally)
	$(call log_info,Platform: $(GOOS)/$(GOARCH))
	$(call log_info,Version: $(VERSION))
	$(call log_info,Installing to $(TF_PLUGIN_DIR)...)
	@mkdir -p $(TF_PLUGIN_DIR)
	@cp $(BINARY) $(TF_PLUGIN_DIR)/$(BINARY)
	$(call log_success,Provider installed successfully!)
	@printf "\n$(COLOR_YELLOW)Note:$(COLOR_RESET) Add this to your ~/.terraformrc:\n\n"
	@printf "$(COLOR_DIM)provider_installation {\n"
	@printf "  dev_overrides {\n"
	@printf "    \"$(HOSTNAME)/$(NAMESPACE)/$(NAME)\" = \"$(TF_PLUGIN_DIR)\"\n"
	@printf "  }\n"
	@printf "  direct {}\n"
	@printf "}$(COLOR_RESET)\n\n"

.PHONY: uninstall
uninstall: ## Remove locally installed provider
	$(call log_header,Uninstalling Provider)
	$(call log_info,Removing $(TF_PLUGIN_DIR)...)
	@rm -rf ~/.terraform.d/plugins/$(HOSTNAME)/$(NAMESPACE)/$(NAME)
	$(call log_success,Provider uninstalled)

.PHONY: clean
clean: ## Clean build artifacts and caches
	$(call log_header,Cleaning)
	$(call log_info,Removing build artifacts...)
	@rm -f $(BINARY)
	@rm -rf dist/
	$(call log_info,Cleaning Go cache...)
	@$(GO) clean -cache -testcache
	$(call log_success,Clean complete)

.PHONY: clean-plugins
clean-plugins: ## Remove all local Terraform plugin installations for this provider
	$(call log_header,Cleaning All Plugin Versions)
	$(call log_info,Removing all versions from ~/.terraform.d/plugins/$(HOSTNAME)/$(NAMESPACE)/$(NAME)...)
	@rm -rf ~/.terraform.d/plugins/$(HOSTNAME)/$(NAMESPACE)/$(NAME)
	$(call log_success,All plugin versions removed)

# ============================================================================
# Testing Targets
# ============================================================================

##@ Testing

.PHONY: test
test: ## Run unit tests
	$(call log_header,Running Unit Tests)
	$(call log_info,Running tests with timeout $(UNIT_TEST_TIMEOUT)...)
	@$(GO) test $(TEST) -v $(TESTARGS) -timeout=$(UNIT_TEST_TIMEOUT)
	$(call log_success,Unit tests passed)

.PHONY: testacc
testacc: ## Run acceptance tests (requires HARNESS_* env vars)
	$(call log_header,Running Acceptance Tests)
	$(call log_warn,This will create real resources in your Harness account)
	$(call log_info,Timeout: $(ACC_TEST_TIMEOUT))
	@TF_ACC=1 $(GO) test $(TEST) -v $(TESTARGS) -timeout $(ACC_TEST_TIMEOUT)
	$(call log_success,Acceptance tests passed)

.PHONY: test-coverage
test-coverage: ## Run tests with coverage report
	$(call log_header,Running Tests with Coverage)
	$(call log_info,Generating coverage report...)
	@$(GO) test $(TEST) -v -coverprofile=coverage.out -covermode=atomic $(TESTARGS)
	@$(GO) tool cover -html=coverage.out -o coverage.html
	$(call log_success,Coverage report generated: coverage.html)
	@$(GO) tool cover -func=coverage.out | tail -1

.PHONY: sweep
sweep: ## Run sweepers to clean up test resources (DANGEROUS)
	$(call log_header,Running Sweepers)
	@printf "$(COLOR_RED)$(COLOR_BOLD)WARNING: This will destroy infrastructure!$(COLOR_RESET)\n"
	@printf "$(COLOR_YELLOW)Only use this in development/test accounts.$(COLOR_RESET)\n\n"
	@read -p "Are you sure? [y/N] " confirm && [ "$$confirm" = "y" ] || exit 1
	$(call log_info,Running sweepers...)
	@$(GO) test $(SWEEP_DIR) -v -sweep=$(SWEEP_RUN) $(SWEEPARGS) -timeout 60m

# ============================================================================
# Code Quality Targets
# ============================================================================

##@ Code Quality

.PHONY: fmt
fmt: ## Format Go source code
	$(call log_header,Formatting Code)
	$(call log_info,Running gofmt...)
	@$(GO) fmt ./...
	$(call log_success,Code formatted)

.PHONY: fmt-check
fmt-check: ## Check if code is formatted
	$(call log_header,Checking Code Format)
	@if [ -n "$$(gofmt -l .)" ]; then \
		printf "$(COLOR_RED)✗$(COLOR_RESET) The following files need formatting:\n"; \
		gofmt -l .; \
		exit 1; \
	fi
	$(call log_success,All files are properly formatted)

.PHONY: vet
vet: ## Run go vet
	$(call log_header,Running Go Vet)
	$(call log_info,Analyzing code...)
	@$(GO) vet ./...
	$(call log_success,No issues found)

.PHONY: lint
lint: ## Run golangci-lint
	$(call log_header,Running Linter)
	@if [ ! -f $(GOLANGCI_LINT) ]; then \
		printf "$(COLOR_YELLOW)⚠$(COLOR_RESET) golangci-lint not found, installing...\n"; \
		$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	$(call log_info,Running golangci-lint...)
	@$(GOLANGCI_LINT) run ./...
	$(call log_success,Linting passed)

.PHONY: check
check: fmt-check vet ## Run all code quality checks
	$(call log_success,All checks passed)

# ============================================================================
# Documentation Targets
# ============================================================================

##@ Documentation

.PHONY: docs
docs: ## Generate provider documentation (preserves subcategories)
	$(call log_header,Generating Documentation)
	@./scripts/generate-docs.sh

.PHONY: docs-validate
docs-validate: ## Validate provider documentation
	$(call log_header,Validating Documentation)
	@if [ ! -f $(TFPLUGINDOCS) ]; then \
		printf "$(COLOR_YELLOW)⚠$(COLOR_RESET) tfplugindocs not found, installing...\n"; \
		$(GO) install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest; \
	fi
	$(call log_info,Validating documentation...)
	@$(TFPLUGINDOCS) validate
	$(call log_success,Documentation is valid)

.PHONY: changelog
changelog: ## Generate changelog from .changelog entries
	$(call log_header,Generating Changelog)
	@./scripts/generate-changelog.sh

# ============================================================================
# Development Setup Targets
# ============================================================================

##@ Development

.PHONY: deps
deps: ## Download and tidy Go dependencies
	$(call log_header,Managing Dependencies)
	$(call log_info,Downloading dependencies...)
	@$(GO) mod download
	$(call log_info,Tidying go.mod...)
	@$(GO) mod tidy
	$(call log_success,Dependencies ready)

.PHONY: deps-upgrade
deps-upgrade: ## Upgrade all dependencies
	$(call log_header,Upgrading Dependencies)
	$(call log_info,Upgrading all dependencies...)
	@$(GO) get -u ./...
	@$(GO) mod tidy
	$(call log_success,Dependencies upgraded)

.PHONY: tools
tools: ## Install development tools
	$(call log_header,Installing Development Tools)
	$(call log_info,Installing tfplugindocs...)
	@$(GO) install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest
	$(call log_info,Installing changelog-build...)
	@$(GO) install github.com/hashicorp/go-changelog/cmd/changelog-build@latest
	$(call log_info,Installing golangci-lint...)
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(call log_success,All tools installed)

.PHONY: dev-setup
dev-setup: deps tools ## Set up complete development environment
	$(call log_header,Development Setup Complete)
	$(call log_success,You're ready to develop!)
	@printf "\n$(COLOR_BOLD)Next steps:$(COLOR_RESET)\n"
	@printf "  1. Run $(COLOR_GREEN)make build$(COLOR_RESET) to build the provider\n"
	@printf "  2. Run $(COLOR_GREEN)make install$(COLOR_RESET) to install locally\n"
	@printf "  3. Run $(COLOR_GREEN)make test$(COLOR_RESET) to run tests\n\n"

.PHONY: version
version: ## Show version information
	$(call log_header,Version Information)
	@printf "  $(COLOR_BOLD)Provider:$(COLOR_RESET)  $(NAME)\n"
	@printf "  $(COLOR_BOLD)Version:$(COLOR_RESET)   $(VERSION)\n"
	@printf "  $(COLOR_BOLD)Go:$(COLOR_RESET)        $(shell $(GO) version | cut -d' ' -f3)\n"
	@printf "  $(COLOR_BOLD)OS/Arch:$(COLOR_RESET)   $(GOOS)/$(GOARCH)\n"
	@printf "  $(COLOR_BOLD)Git:$(COLOR_RESET)       $(shell git describe --tags --always --dirty 2>/dev/null || echo 'N/A')\n"
	@printf "\n"

