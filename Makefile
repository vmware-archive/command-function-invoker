.PHONY: build clean dockerize debug-dockerize gen-mocks test gen-proto docs verify-docs
OUTPUT = command-function-invoker

GO_SOURCES = $(shell find cmd pkg -type f -name '*.go')
TAG ?= $(shell cat VERSION)

build: $(OUTPUT)

docs:
	RIFF_INVOKER_PATHS=command-invoker.yaml riff docs -d docs -c "init command"
	RIFF_INVOKER_PATHS=command-invoker.yaml riff docs -d docs -c "create command"
	$(call embed_readme,init)
	$(call embed_readme,create)

define embed_readme
    $(shell cat README.md | perl -e 'open(my $$fh, "docs/riff_$(1)_command.md"); my $$doc = join("", <$$fh>) =~ s/^#/##/rmg; print join("", <STDIN>) =~ s/(?<=<!-- riff-$(1) -->\n).*(?=\n<!-- \/riff-$(1) -->)/\n$$doc/sr' > README.$(1).md; mv README.$(1).md README.md)
endef

verify-docs: docs
	git diff --exit-code -- docs
	git diff --exit-code -- README.md

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

gen-proto:
	protoc -I $(FN_PROTO_PATH)/ $(FN_PROTO_PATH)/function.proto --go_out=plugins=grpc:pkg/function

clean:
	rm -f $(OUTPUT)

dockerize: $(GO_SOURCES) vendor
	docker build . -t "projectriff/command-function-invoker:latest"
	docker tag "projectriff/command-function-invoker:latest" "projectriff/command-function-invoker:$(TAG)"

debug-dockerize: $(GO_SOURCES) vendor
	docker build . -t "projectriff/command-function-invoker:latest" -f Dockerfile-debug
	docker tag "projectriff/command-function-invoker:latest" "projectriff/command-function-invoker:$(TAG)"
