# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	@mkdir -p $(LOCALBIN)

TEST_ARGS ?= ""

## Tool Binaries
GINKGO ?= $(LOCALBIN)/ginkgo
GOLANGCILINT ?= $(LOCALBIN)/golangci-lint

# Image URL to use all building/pushing image targets
IMG ?= doc-gen
TAG ?= dev

.PHONY: all
all: build

.PHONY: build
build:  ## Build doc-gen binary
	go build -o bin/doc-gen main.go

.PHONY: test
test:
	$(GINKGO) $(TEST_ARGS) tests/...

.PHONY: lint
lint: lint-commits lint-sources## Run all linters and fail on error

.PHONY: lint-sources
lint-sources: ## Run golangci-lint and fail on error
	@echo "Linting go sources ..."
	@$(GOLANGCILINT) --concurrency 2 run ./...

.PHONY: lint-fix
lint-fix: ## Fix whatever golangci-lint can fix
	@$(GOLANGCILINT) run ./... --fix

.PHONY: lint-commits
lint-commits:  ## Run commitlint and fail on error
	@echo "Linting commits ..."
	@npm i -g @commitlint/config-conventional @commitlint/cli
	@commitlint -x @commitlint/config-conventional --edit

.PHONY: clean-tools
clean-tools: $(LOCALBIN) ## Cleans (delete) all binary tools
	@echo "Cleaning tools"
	@find $(LOCALBIN) -type f -delete

.PHONY: download
download: ## Download all project dependencies
	@echo "Downloading go.mod dependencies"
	@go mod download

.PHONY: install-tools
install-tools: download ## Installs all required GO tools
	@echo "Installing GO tools"
	@cat build/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -I % sh -c 'GOBIN=$(LOCALBIN) go install %'

.PHONY: reinstall-tools
re-install-tools: clean-tools install-tools ## Clean and install tools again

.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	docker build -t ${IMG}:${TAG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMG}:${TAG}

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

