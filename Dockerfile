FROM golang:1.15.8-buster

RUN go get github.com/markbates/refresh

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_USER app
ENV APP_HOME /go/standard-deviation/src

ARG GROUP_ID
ARG USER_ID

RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

USER $APP_USER
WORKDIR $APP_HOME

EXPOSE 8010

CMD ["refresh", "run", "-c", "../refresh.yml"]
