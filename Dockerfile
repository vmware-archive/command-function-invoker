FROM golang:1.9 as builder

ARG PACKAGE=github.com/projectriff/command-function-invoker
ARG COMMAND=command-function-invoker

WORKDIR /go/src/${PACKAGE}
COPY vendor/ vendor/
COPY cmd/ cmd/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo cmd/${COMMAND}.go

###########

FROM alpine

ARG PACKAGE=github.com/projectriff/command-function-invoker
ARG COMMAND=command-function-invoker

COPY --from=builder /go/src/${PACKAGE}/${COMMAND} /${COMMAND}

CMD ["/command-function-invoker"]