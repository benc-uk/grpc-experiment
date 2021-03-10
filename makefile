# Example and common variables
VERSION := 0.0.1
BUILD_INFO := Manual build 
SERVER_DIR := ./server
CLIENT_GO_DIR := ./client-go

# Most likely want to override these when calling `make image`
IMAGE_REG ?= ghcr.io
IMAGE_REPO ?= benc-uk/grpc-echoserver
IMAGE_TAG ?= latest
IMAGE_PREFIX := $(IMAGE_REG)/$(IMAGE_REPO)

.DEFAULT_GOAL := help

help:  ## This help message 😃
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

lint:  ## Lint & format, will not fix but sets exit code on error
	go get github.com/golangci/golangci-lint/cmd/golangci-lint; golangci-lint run $(SERVER_DIR)/... --fix 

lint-fix:  ## Lint & format, will try to fix errors and modify code
	go get github.com/golangci/golangci-lint/cmd/golangci-lint && golangci-lint run $(SERVER_DIR)/... --fix 

image:  ## Build container image from Dockerfile
	docker build --file ./build/Dockerfile \
	--build-arg BUILD_INFO="$(BUILD_INFO)" \
	--build-arg VERSION="$(VERSION)" \
	--tag $(IMAGE_PREFIX):$(IMAGE_TAG) . 

push:  ## Push container image to registry
	docker push $(IMAGE_PREFIX):$(IMAGE_TAG)

build-server: proto-go  ## Run a build of server binary
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o ./bin/server $(SERVER_DIR)/...

run-server: proto-go  ## Run server locally, with hot reload
	air -c .air.toml

build-client-go: proto-go  ## Run a build of Go client binary
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o ./bin/client $(CLIENT_GO_DIR)/...

run-client-go: proto-go  ## Run Go client locally, with hot reload
	go run ./client-go/main.go

proto-go:  ## Compile API bindings for Go
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./api/hello.proto

proto-python:  .venv  ## Compile API bindings for Python
	python -m grpc_tools.protoc --proto_path=./api  --python_out=./api --grpc_python_out=./api hello.proto

clean:  ## Tidy up!
	@rm -rf .venv
	@rm -rf bin/server
	@rm -rf bin/client
	@rm -rf api/*.go; rm -rf api/*.py; rm -rf api/__pycache__

.venv: .venv/touchfile

.venv/touchfile: requirements.txt
	python3 -m venv .venv
	. .venv/bin/activate; pip install -Ur requirements.txt
	touch .venv/touchfile
