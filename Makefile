
GIT_REPOSITORY_ID ?= definition-driven-development
GIT_USER_ID ?= edenreich

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  generate-sdk-go               Generate the go SDK"
	@echo "  generate-api                  Generate the go API"
	@echo "  generate-protobuf-schema      Generate the protobuf schema"
	@echo "  generate                      Generate the SDK's and API"
	@echo "  regenerate-sdk-go             Re-generate the go SDK"
	@echo "  regenerate-api                Re-generate the go API"
	@echo "  regenerate-protobuf-schema    Re-generate the protobuf schema"
	@echo "  regenerate                    Re-generate the SDK's and API"
	@echo "  tidy-sdk-go                   Tidy the go SDK"
	@echo "  tidy-api                      Tidy the API"
	@echo "  tidy                          Tidy the API and SDK's"
	@echo "  run-api                       Run the API"
	@echo "  test-api                      Test the API"
	@echo "  test-sdk-go                   Test the go SDK"
	@echo "  test                          Test the API and SDK's"
	@echo "  openapi                       Run the openapi-generator-cli"
	@echo "  clean                         Clean the SDK's, proto files and API"
	@echo "  help                          Show this help message"

default: help

generate-sdk-go:
	@echo "Generating go SDK..."
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name sdk \
		-o sdk/go

generate-api:
	@echo "Generating go API..."
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go-server \
		-t templates/go-server \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name api \
		-o api

generate-protobuf-schema:
	@echo "Generating protobuf schema..."
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g protobuf-schema \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name sdk \
		-o protobuf

generate: generate-sdk-go generate-api generate-protobuf-schema

regenerate-sdk-go:
	@echo "Re-generating go SDK..."
	@rm -rf sdk/go
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g go \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
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
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name api \
		-o api

regenerate-protobuf-schema:
	@echo "Re-generating protobuf schema..."
	@rm -rf protobuf
	@docker run --rm \
		-v $(PWD):/local -w /local openapitools/openapi-generator-cli generate \
		-i cats_api.yaml \
		-g protobuf-schema \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name sdk \
		-o protobuf

regenerate: regenerate-sdk-go regenerate-api regenerate-protobuf-schema

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

clean:
	@rm -rf api
	@rm -rf sdk
	@rm -rf protobuf

.PHONY: generate-sdk-go generate-api generate-protobuf-schema generate regenerate-sdk-go regenerate-api regenerate-protobuf-schema regenerate tidy-sdk-go tidy-api tidy run-api test-api test-sdk-go test openapi
