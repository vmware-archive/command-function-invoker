.PHONY: all build clean gen-mocks test release
OUTPUT = command-function-invoker

GO_SOURCES = $(shell find cmd pkg -type f -name '*.go')
TAG ?= $(shell cat VERSION)

all: build test

build: $(OUTPUT)

test: build
	go test -v ./...

$(OUTPUT): $(GO_SOURCES) vendor
	go build cmd/command-function-invoker.go

ifdef CI
vendor:
else
vendor: glide.lock
	glide install -v --force
endif

glide.lock: glide.yaml
	glide up -v --force

gen-mocks: $(GO_SOURCES)
	go get -u github.com/vektra/mockery/.../
	go generate ./...

clean:
	rm -f $(OUTPUT)
	rm -f $(OUTPUT).tgz

release: build LICENSE README.md
	tar cvzf $(OUTPUT).tgz LICENSE README.md $(OUTPUT)
