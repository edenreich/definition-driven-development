
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  generate-sdk-go     Generate go client"
	@echo "  generate-api        Generate go server"
	@echo "  tidy-sdk-go         Tidy go SDK"
	@echo "  tidy-api            Tidy API"
	@echo "  tidy                Tidy API and SDK's"
	@echo "  generate            Generate SDK's and API"

default: help

generate-sdk-go:
	@echo "Generating go client..."
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go \
		--git-repo-id definition-driven-development \
		--git-user-id edenreich \
		-o sdk/go

generate-api:
	@echo "Generating go server..."
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go-server \
		-t templates/go-server \
		--git-repo-id definition-driven-development \
		--git-user-id edenreich \
		-o api

generate: generate-sdk-go generate-api

tidy-sdk-go:
	@echo "Tidying go client..."
	@cd sdk/go && go mod tidy

tidy-api:
	@echo "Tidying go server..."
	@cd api && go mod tidy

tidy: tidy-sdk-go tidy-api

openapi:
	@docker run --rm -v $(PWD):/local -w /local openapitools/openapi-generator-cli $(ARGS)

.PHONY: generate-sdk-go generate-api generate tidy-sdk-go tidy-api tidy openapi
