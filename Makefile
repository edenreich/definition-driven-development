
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  generate-sdk-go     Generate the go SDK"
	@echo "  generate-api        Generate the go API"
	@echo "  generate            Generate the SDK's and API"
	@echo "  regenerate-sdk-go   Re-generate the go SDK"
	@echo "  regenerate-api      Re-generate the go API"
	@echo "  regenerate          Re-generate the SDK's and API"
	@echo "  tidy-sdk-go         Tidy the go SDK"
	@echo "  tidy-api            Tidy the API"
	@echo "  tidy                Tidy the API and SDK's"
	@echo "  run-api             Run the API"
	@echo "  test-api            Test the API"
	@echo "  test-sdk-go         Test the go SDK"
	@echo "  test                Test the API and SDK's"
	@echo "  openapi             Run the openapi-generator-cli"
	@echo "  help                Show this help message"

default: help

generate-sdk-go:
	@echo "Generating go SDK..."
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go \
		--git-repo-id definition-driven-development \
		--git-user-id edenreich \
		--package-name sdk \
		-o sdk/go

generate-api:
	@echo "Generating go API..."
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go-server \
		-t templates/go-server \
		--git-repo-id definition-driven-development \
		--git-user-id edenreich \
		--package-name api \
		-o api

generate: generate-sdk-go generate-api

regenerate-sdk-go:
	@echo "Re-generating go SDK..."
	@rm -rf sdk/go
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go \
		--git-repo-id definition-driven-development \
		--git-user-id edenreich \
		--package-name sdk \
		-o sdk/go

regenerate-api:
	@echo "Re-generating go API..."
	@rm -rf api
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go-server \
		-t templates/go-server \
		--git-repo-id definition-driven-development \
		--git-user-id edenreich \
		--package-name api \
		-o api

regenerate: regenerate-sdk-go regenerate-api

tidy-sdk-go:
	@echo "Tidying go SDK..."
	@cd sdk/go && go mod tidy

tidy-api:
	@echo "Tidying go API..."
	@cd api && go mod tidy

tidy: tidy-sdk-go tidy-api

run-api:
	@echo "Running go API..."
	@cd api && go run main.go

test-api:
	@echo "Testing go API..."
	@cd api && go test ./...

test-sdk-go:
	@echo "Testing go SDK..."
	@cd sdk/go && go test ./...

test: test-sdk-go test-api

openapi:
	@docker run --rm -v $(PWD):/local -w /local openapitools/openapi-generator-cli $(ARGS)

.PHONY: generate-sdk-go generate-api generate regenerate-sdk-go regenerate-api regenerate tidy-sdk-go tidy-api tidy run-api test-api test-sdk-go test openapi
