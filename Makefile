generator_version = v4.2.2
api_doc = https://raw.githubusercontent.com/thepracticaldev/dev.to/7d0aeeefe5cf6250c5b58ae6b631bfe41fe5bf4a/docs/api_v0.yml

.PHONY: gen

gen:
	docker run --rm \
		--user "$$(id -u):$$(id -g)" \
		-v $(PWD):/local openapitools/openapi-generator-cli:$(generator_version) generate \
		-i $(api_doc) \
		-g go \
		-o /local/pkg/devto \
		-c /local/openapi-generator-config.yml && \
		gofmt -w pkg/devto/..

lint:
	docker run --rm -v $$(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v
