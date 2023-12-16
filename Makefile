
OPENAPI_SPEC ?= pets_api.yaml
OPENAPI_GENERATOR_IMAGE ?= openapitools/openapi-generator-cli
GIT_REPOSITORY_ID ?= definition-driven-development
GIT_USER_ID ?= edenreich

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  fetch-templates               Fetching vendors generators templates"
	@echo "  generate-http-sdk-go          Generate the HTTP go SDK"
	@echo "  generate-http-api             Generate the HTTP go API"
	@echo "  generate-protobuf-schema      Generate the protobuf schema"
	@echo "  generate                      Generate the HTTP SDK's and API"
	@echo "  regenerate-http-sdk-go        Re-generate the HTTP go SDK"
	@echo "  regenerate-http-api           Re-generate the HTTP go API"
	@echo "  regenerate-protobuf-schema    Re-generate the protobuf schema"
	@echo "  regenerate                    Re-generate the HTTP SDK's and API"
	@echo "  tidy-http-sdk-go              Tidy the HTTP go SDK"
	@echo "  tidy-http-api                 Tidy the HTTP API"
	@echo "  tidy                          Tidy the HTTP API and SDK's"
	@echo "  run-http-api                  Run the HTTP API"
	@echo "  test-http-api                 Test the HTTP API"
	@echo "  test-http-sdk-go              Test the HTTP go SDK"
	@echo "  test                          Test the HTTP API and SDK's"
	@echo "  openapi                       Run the openapi-generator-cli"
	@echo "  clean                         Clean the SDK's, proto files and API"
	@echo "  help                          Show this help message"

default: help

fetch-templates:
	@echo "Fetching vendors generators templates..."
	@mkdir -p templates
	@make openapi ARGS='author template -g go -o templates/go'
	@make openapi ARGS='author template -g go-server -o templates/go-server'
	@make openapi ARGS='author template -g protobuf-schema -o templates/protobuf-schema'

generate-http-sdk-go:
	@echo "Generating go HTTP SDK..."
	@docker run --rm \
		-v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) generate \
		-i $(OPENAPI_SPEC) \
		-g go \
		-t templates/go \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name sdk \
		-o http/sdk/go

generate-http-api:
	@echo "Generating go HTTP API..."
	@docker run --rm \
		-v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) generate \
		-i $(OPENAPI_SPEC) \
		-g go-server \
		-t templates/go-server \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name api \
		-o http/api

generate-protobuf-schema:
	@echo "Generating protobuf schema..."
	@docker run --rm \
		-v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) generate \
		-i $(OPENAPI_SPEC) \
		-g protobuf-schema \
		-t templates/protobuf-schema \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		-o grpc/protobuf

generate: generate-http-sdk-go generate-http-api generate-protobuf-schema

regenerate-http-sdk-go:
	@echo "Re-generating go HTTP SDK..."
	@rm -rf sdk/go
	@docker run --rm \
		-v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) generate \
		-i $(OPENAPI_SPEC) \
		-g go \
		-t templates/go \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name sdk \
		-o http/sdk/go

regenerate-http-api:
	@echo "Re-generating go API..."
	@rm -rf api
	@docker run --rm \
		-v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) generate \
		-i $(OPENAPI_SPEC) \
		-g go-server \
		-t templates/go-server \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name api \
		-o http/api

regenerate-protobuf-schema:
	@echo "Re-generating protobuf schema..."
	@rm -rf protobuf
	@docker run --rm \
		-v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) generate \
		-i $(OPENAPI_SPEC) \
		-g protobuf-schema \
		-t templates/protobuf-schema \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		-o grpc/protobuf

regenerate: regenerate-http-sdk-go regenerate-http-api regenerate-protobuf-schema

tidy-http-sdk-go:
	@echo "Tidying go HTTP SDK..."
	@cd http/sdk/go && go mod tidy

tidy-http-api:
	@echo "Tidying go HTTP API..."
	@cd http/api && go mod tidy

tidy: tidy-http-sdk-go tidy-http-api

run-http-api:
	@echo "Running go HTTP API..."
	@cd http/api && go run main.go

test-http-api:
	@echo "Testing go API..."
	@cd http/api && go test ./...

test-http-sdk-go:
	@echo "Testing go SDK..."
	@cd http/sdk/go && go test ./...

test: test-http-sdk-go test-http-api

openapi:
	@docker run --rm -v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) $(ARGS)

clean:
	@rm -rf grpc
	@rm -rf http
	@rm -rf templates

.PHONY: fetch-templates generate-http-sdk-go generate-http-api generate-protobuf-schema generate regenerate-http-sdk-go regenerate-http-api regenerate-protobuf-schema regenerate tidy-http-sdk-go tidy-http-api tidy run-http-api test-http-api test-http-sdk-go test openapi
