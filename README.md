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

# Configuration

Environmental variables

| Setting / Variable | Purpose                         | Default |
| ------------------ | ------------------------------- | ------- |
| PORT               | Port the server will listen on. | 8000    |

# Repository Structure

A brief description of the top-level directories of this project is as follows:

```
/api        - Details of the API specification & docs
/build      - Build configuration e.g. Dockerfiles
/charts     - Helm charts
/deploy     - Deployment and infrastructure as code, inc Kubernetes
/scripts    - Bash and other supporting scripts
/src        - Source code
/test       - Testing, mock data and API + load tests
```

# API

See the [API documentation](./api/) for full infomration about the API(s).  
Optional. Delete this section if project as no API.

# Known Issues

List any known bugs or gotchas.

# Change Log

See [complete change log](./CHANGELOG.md)

# License

This project uses the MIT software license. See [full license file](./LICENSE)

# Acknowledgements

Optional. Put acknowledgements and credits here, if any
