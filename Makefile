.PHONY: build clean dockerize debug-dockerize gen-mocks test gen-proto
OUTPUT = shell-function-invoker

GO_SOURCES = $(shell find cmd pkg -type f -name '*.go')
TAG ?= $(shell cat VERSION)

build: $(OUTPUT)

test: build
	go test -v ./...

$(OUTPUT): $(GO_SOURCES) vendor
	go build cmd/shell-function-invoker.go

vendor: glide.lock
	glide install -v --force

glide.lock: glide.yaml
	glide up -v --force

gen-mocks: $(GO_SOURCES)
	go get -u github.com/vektra/mockery/.../
	go generate ./...

gen-proto:
	protoc -I $(FN_PROTO_PATH)/ $(FN_PROTO_PATH)/function.proto --go_out=plugins=grpc:pkg/function

clean:
	rm -f $(OUTPUT)

dockerize: $(GO_SOURCES) vendor
	docker build . -t "projectriff/shell-function-invoker:latest"
	docker tag "projectriff/shell-function-invoker:latest" "projectriff/shell-function-invoker:$(TAG)"

debug-dockerize: $(GO_SOURCES) vendor
	docker build . -t "projectriff/shell-function-invoker:latest" -f Dockerfile-debug
	docker tag "projectriff/shell-function-invoker:latest" "projectriff/shell-function-invoker:$(TAG)"
