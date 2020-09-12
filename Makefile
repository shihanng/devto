LDFLAGS := -ldflags="-s -w"

generator_version = v4.2.2
api_doc = https://raw.githubusercontent.com/thepracticaldev/dev.to/7d0aeeefe5cf6250c5b58ae6b631bfe41fe5bf4a/docs/api_v0.yml

.PHONY: test gen lint mod-check install

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	go test -race -covermode atomic -coverprofile=profile.cov -v ./... -count=1

gen:
	docker run --rm \
		--user "$$(id -u):$$(id -g)" \
		-v $(PWD):/local openapitools/openapi-generator-cli:$(generator_version) generate \
		-i $(api_doc) \
		-g go \
		-o /local/pkg/devto \
		-c /local/openapi-generator-config.yml && \
		gofmt -w pkg/devto/..

lint: ## Run GolangCI-Lint
	docker run --rm -v $$(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v

mod-check: ## Run check on go mod tidy
	go mod tidy && git --no-pager diff --exit-code -- go.mod go.sum

devto: ## Build devto binary
	go build $(LDFLAGS) -o devto

install: ## Install devto into $GOBIN
	go install $(LDFLAGS)

clean: ## Cleanup
	rm devto
