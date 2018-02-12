FROM golang:1.9 as builder

ARG PACKAGE=github.com/projectriff/shell-function-invoker
ARG COMMAND=shell-function-invoker

WORKDIR /go/src/${PACKAGE}
COPY vendor/ vendor/
COPY cmd/ cmd/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo cmd/${COMMAND}.go

###########

FROM alpine

ARG PACKAGE=github.com/projectriff/shell-function-invoker
COPY --from=builder /go/src/${PACKAGE}/shell-function-invoker /shell-function-invoker

CMD ["/shell-function-invoker"]