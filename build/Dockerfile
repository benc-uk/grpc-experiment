FROM golang:1.18-alpine as build
WORKDIR /build
ARG VERSION="0.0.0"
ARG BUILD_INFO="Not set"

# Install system dependencies
RUN apk update && apk add git gcc musl-dev make protoc

# Fetch and cache Go modules
COPY go.mod .
COPY go.sum .
COPY makefile .
RUN go mod download
# Plugins needed by protoc
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Copy in Go source files
COPY server/ ./server
COPY api/ ./api
RUN mkdir bin

# Now run the build
RUN make build-server

# ================================================================================================

FROM scratch
WORKDIR /app 

# Copy the Go server binary
COPY --from=build /build/bin/server . 
ENV PORT 50051

EXPOSE 50051
CMD [ "./server" ]
