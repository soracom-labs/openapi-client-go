GENERATED_DIR ?= openapi
REMOVE_GENERATED_FILES ?= .openapi-generator api docs go.mod go.sum README.md git_push.sh .travis.yml .openapi-generator-ignore .gitignore
SOURCE_FILES ?= $(shell find . -not -path "./$(GENERATED_DIR)/*" -type f -name '*.go' -print)
INPUT_SPECS ?= api sandbox

GIT_REVISION ?= $(shell git rev-parse --short HEAD)
GIT_TAG ?= $(shell git describe --tags || echo "v0.0.0")

GOLANGCI_LINT_VERSION ?= 1.43.0
GOLANGCI_LINT ?= golangci-lint

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

.PHONY: install-deps-dev
install-deps-dev: ## install dependencies for development
	# https://golangci-lint.run/usage/install/#linux-and-windows
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v$(GOLANGCI_LINT_VERSION)

.PHONY: format
format: ## format codes
	gofmt -s -w $(SOURCE_FILES)

.PHONY: lint
lint: ## lint
	$(GOLANGCI_LINT) run -v

.PHONY: test
test: ## run tests
	SORACOM_SIM_ID=$(SORACOM_SIM_ID) \
	SORACOM_AUTH_KEY=$(SORACOM_AUTH_KEY) \
	SORACOM_AUTH_KEY_ID=$(SORACOM_AUTH_KEY_ID) \
	go test -cover -v ./... -count=1

.PHONY: install-openapi-generator-cli
install-openapi-generator-cli: ## install openapi-generator-cli
	which openapi-generator-cli || npm install -g @openapitools/openapi-generator-cli

.PHONY: generate
generate: ## run OpenAPI Generator
	@for spec in $(INPUT_SPECS) ; do \
		make generate-spec INPUT_SPEC=$$spec ; \
    done

.PHONY: generate-spec
generate-spec: ## run OpenAPI Generator
	npx @openapitools/openapi-generator-cli generate \
		--input-spec specs/$(INPUT_SPEC).yaml \
		--generator-name go \
		--output $(GENERATED_DIR)/$(INPUT_SPEC) \
		--package-name $(INPUT_SPEC) \
		--git-host github.com \
		--git-user-id soracom-labs \
		--git-repo-id openapi-client-go
	cd $(GENERATED_DIR)/$(INPUT_SPEC) && rm -rf $(REMOVE_GENERATED_FILES)

.PHONY: generate-diff-check
generate-diff-check: ## check if generated codes are the same as tracked ones
	git diff --exit-code --relative $(GENERATED_DIR)

.PHONY: ci-test
ci-test: install-openapi-generator-cli generate generate-diff-check lint test ## ci test
