FROM golang:1.15.8

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_HOME /go/standard-deviation

ADD . $APP_HOME

WORKDIR $APP_HOME

RUN go mod vendor

EXPOSE 8080

CMD ["go", "run", "src/main.go"]
