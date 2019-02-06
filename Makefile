.PHONY: all build clean gen-mocks test release
OUTPUT = command-function-invoker

GO_SOURCES = $(shell find cmd pkg -type f -name '*.go')
TAG ?= $(shell cat VERSION)

all: build test

build: $(OUTPUT)

test:
	GO111MODULE=on go test -v ./...

$(OUTPUT): $(GO_SOURCES)
	GO111MODULE=on go build cmd/command-function-invoker.go

gen-mocks: $(GO_SOURCES)
	GO111MODULE=off go get -u github.com/vektra/mockery/.../
	GO111MODULE=on go generate ./...

clean:
	rm -f $(OUTPUT)
	rm -f $(OUTPUT).tgz

release: build LICENSE README.md
	tar cvzf $(OUTPUT).tgz LICENSE README.md $(OUTPUT)
