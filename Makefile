
OPENAPI_SPEC ?= pets_api.yaml
OPENAPI_GENERATOR_IMAGE ?= openapitools/openapi-generator-cli
GIT_REPOSITORY_ID ?= definition-driven-development
GIT_USER_ID ?= edenreich

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  fetch-templates               Fetching vendors generators templates"
	@echo "  generate-http-sdk-go          Generate the go HTTP SDK"
	@echo "  generate-http-api             Generate the go HTTP API"
	@echo "  generate-protobuf-schema      Generate the protobuf schema"
	@echo "  generate-grpc-go              Generate the go gRPC"
	@echo "  generate                      Generate the HTTP SDK's and API"
	@echo "  regenerate-http-sdk-go        Re-generate the go HTTP SDK"
	@echo "  regenerate-http-api           Re-generate the go HTTP API"
	@echo "  regenerate-protobuf-schema    Re-generate the protobuf schema"
	@echo "  regenerate-grpc-go            Re-generate the go gRPC"
	@echo "  regenerate                    Re-generate the HTTP SDK's and API"
	@echo "  tidy-http-sdk-go              Tidy the go HTTP SDK"
	@echo "  tidy-http-api                 Tidy the go HTTP API"
	@echo "  tidy                          Tidy the go HTTP API and SDK's"
	@echo "  run-http-api                  Run the HTTP API"
	@echo "  test-http-api                 Test the HTTP API"
	@echo "  test-http-sdk-go              Test the go HTTP SDK"
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
	@echo "Running go fmt..."
	@cd http/sdk/go && go fmt ./...
	@echo "Running go mod tidy..."
	@cd http/sdk/go && go mod tidy
	@echo "Remove unused imports..."
	@docker run --rm -v $(PWD)/http/sdk/go:/app -w /app golang:1.21 bash -c 'go install golang.org/x/tools/cmd/goimports@latest && goimports -w .'
	@echo "Running golangci-lint fix..."
	@docker run --rm -v $(PWD)/http/sdk/go:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v --fix

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
	@echo "Running go fmt..."
	@cd http/api && go fmt ./...
	@echo "Running go mod tidy..."
	@cd http/api && go mod tidy
	@echo "Remove unused imports..."
	@docker run --rm -v $(PWD)/http/api:/app -w /app golang:1.21 bash -c 'go install golang.org/x/tools/cmd/goimports@latest && goimports -w .'
	@echo "Running golangci-lint fix..."
	@docker run --rm -v $(PWD)/http/api:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v --fix

generate-protobuf-schema:
	@echo "Generating protobuf schema..."
	@docker run --rm \
		-v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) generate \
		-i $(OPENAPI_SPEC) \
		-g protobuf-schema \
		-t templates/protobuf-schema \
		--git-repo-id $(GIT_REPOSITORY_ID) \
		--git-user-id $(GIT_USER_ID) \
		--package-name grpc \
		--additional-properties=sourceFolder=grpc \
		-o grpc/protobuf

generate-grpc-go: generate-protobuf-schema
	@echo "Generating go gRPC..."
	@docker run --rm \
		-v $(PWD)/grpc:/grpc -w /grpc/protobuf golang:1.21 \
		bash -c 'apt update && apt install -y protobuf-compiler && go install github.com/golang/protobuf/protoc-gen-go@latest && \
			protoc --go_out=plugins=grpc,paths=source_relative:../ models/*.proto && \
			protoc --go_out=plugins=grpc,paths=source_relative:../ services/*.proto'

generate: generate-http-sdk-go generate-http-api generate-protobuf-schema generate-grpc-go

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
	@echo "Running go fmt..."
	@cd http/sdk/go && go fmt ./...
	@echo "Running go mod tidy..."
	@cd http/sdk/go && go mod tidy
	@echo "Remove unused imports..."
	@docker run --rm -v $(PWD)/http/sdk/go:/app -w /app golang:1.21 bash -c 'go install golang.org/x/tools/cmd/goimports@latest && goimports -w .'
	@echo "Running golangci-lint fix..."
	@docker run --rm -v $(PWD)/http/sdk/go:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v --fix

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
	@echo "Running go fmt..."
	@cd http/api && go fmt ./...
	@echo "Running go mod tidy..."
	@cd http/api && go mod tidy
	@echo "Remove unused imports..."
	@docker run --rm -v $(PWD)/http/api:/app -w /app golang:1.21 bash -c 'go install golang.org/x/tools/cmd/goimports@latest && goimports -w .'
	@echo "Running golangci-lint fix..."
	@docker run --rm -v $(PWD)/http/api:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v --fix

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

regenerate-grpc-go: generate-protobuf-schema
	@echo "Re-generating go gRPC..."
	@rm -rf grpc
	@mkdir -p grpc
	@docker run --rm \
		-v $(PWD)/grpc:/grpc -w /grpc/protobuf golang:1.21 \
		bash -c 'apt update && apt install -y protobuf-compiler && go install github.com/golang/protobuf/protoc-gen-go@latest && \
			protoc --go_out=plugins=grpc,paths=source_relative:../ models/*.proto && \
			protoc --go_out=plugins=grpc,paths=source_relative:../ services/*.proto'

regenerate: regenerate-http-sdk-go regenerate-http-api regenerate-protobuf-schema regenerate-grpc-go

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
	@echo "Testing go HTTP API..."
	@cd http/api && go test ./...

test-http-sdk-go:
	@echo "Testing go HTTP SDK..."
	@cd http/sdk/go && go test ./...

test: test-http-sdk-go test-http-api

openapi:
	@docker run --rm -v $(PWD):/local -w /local $(OPENAPI_GENERATOR_IMAGE) $(ARGS)

clean:
	@rm -rf grpc
	@rm -rf http
	@rm -rf templates

.PHONY: fetch-templates generate-http-sdk-go generate-http-api generate-protobuf-schema generate-grpc-go generate regenerate-http-sdk-go regenerate-http-api regenerate-protobuf-schema regenerate-grpc-go regenerate tidy-http-sdk-go tidy-http-api tidy run-http-api test-http-api test-http-sdk-go test openapi
