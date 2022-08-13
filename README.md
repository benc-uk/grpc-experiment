# gRPC Experiment

Purpose and description of this project

Goals:

- Learn gRPC

Use cases & key features:

- Very little

Supporting technologies and libraries:

- Go gRPC
- Protobuf

# Getting Started

## Installing / Deploying

- If the project can be installed (such as a command line tool or library)
- Or deployed to Kubernetes, public cloud etc

## Running as container

Notes on running the project from Docker image / container

## Running locally

### Install Pre-reqs

Install `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc` and `air`

```
sudo apt install protobuf-compiler
go get google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u github.com/cosmtrek/air
```

Python prereq

```
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```
# Use Make

Make is the "frontend" for working locally with this repo

```text
$ make
build-client-go      Run a build of Go client binary, into ./bin
build-server         Run a build of server binary, into ./bin
clean                Tidy up!
help                 This help message 😃
lint-fix             Lint & format, will try to fix errors and modify code
lint                 Lint & format, will not fix but sets exit code on error
proto-go             Compile API bindings for Go, into ./api
proto-python         Compile API bindings for Python, into ./api
run-client-go        Run Go client locally, with hot reload
run-server           Run server locally, with hot reload
```

# Repository Structure

A brief description of the top-level directories of this project is as follows:

```text
📂
├── api            - Proto definition and holds auto-gen code
├── build          - Dockerfiles
├── client-go      - Client written in Go
├── client-python  - Client written in Python
├── deploy         - Some Kubernetes manifests
└── server         - Server in Go
```

# API

See the [API documentation](./api/) for full infomration about the API(s).  
Optional. Delete this section if project as no API.
