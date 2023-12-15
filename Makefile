
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  generate-go-client  Generate go client"
	@echo "  generate-go-server  Generate go server"
	@echo "  generate            Generate go client and server"

default: help

generate-go-client:
	@echo "Generating go client..."
	@docker run --rm \
		-v $(PWD):/local openapitools/openapi-generator-cli generate \
		-i /local/cats_api.yaml \
		-g go \
		-o /local/src/go/client

generate-go-server:
	@echo "Generating go server..."
	@docker run --rm \
		-v $(PWD):/local openapitools/openapi-generator-cli generate \
		-i /local/cats_api.yaml \
		-g go-server \
		-o /local/src/go/server

generate: generate-go-client generate-go-server

openapi:
	@docker run --rm -v $(PWD):/local openapitools/openapi-generator-cli

.PHONY: generate-client generate-server generate openapi
