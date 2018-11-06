.PHONY: build clean dockerize debug-dockerize gen-mocks test release
OUTPUT = command-function-invoker

GO_SOURCES = $(shell find cmd pkg -type f -name '*.go')
TAG ?= $(shell cat VERSION)

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

dockerize: $(GO_SOURCES) vendor
	docker build . -t "projectriff/command-function-invoker:latest"
	docker tag "projectriff/command-function-invoker:latest" "projectriff/command-function-invoker:$(TAG)"

debug-dockerize: $(GO_SOURCES) vendor
	docker build . -t "projectriff/command-function-invoker:latest" -f Dockerfile-debug
	docker tag "projectriff/command-function-invoker:latest" "projectriff/command-function-invoker:$(TAG)"

release: build LICENSE README.md
	tar cvzf $(OUTPUT).tgz LICENSE README.md $(OUTPUT)
