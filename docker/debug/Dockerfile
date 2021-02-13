FROM golang:1.15.8-alpine3.13 AS build-env

ENV CGO_ENABLED 0

RUN apk add --no-cache git

WORKDIR /go/standard-deviation/
ADD . /go/standard-deviation/

RUN go build -o /go/standard-deviation/srv ./src

WORKDIR /go/src/
RUN go get github.com/go-delve/delve/cmd/dlv

FROM alpine:3.8

WORKDIR /
COPY --from=build-env /go/standard-deviation/srv /
COPY --from=build-env /go/bin/dlv /

EXPOSE 8080 40000

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/srv"]