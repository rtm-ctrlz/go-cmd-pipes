FROM golang:alpine as builder

COPY . /go/src

WORKDIR /go/src

RUN set -ex; \
    go build -a -tags netgo -ldflags '-w -extldflags "-static"'

FROM scratch
ENTRYPOINT ["/app"]
COPY --from=builder /go/src/go-cmd-pipes /app
