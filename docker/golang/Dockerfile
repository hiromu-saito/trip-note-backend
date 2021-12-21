FROM golang:1.17.5-alpine3.15


RUN apk update && \
  apk upgrade &&\
  apk add git

WORKDIR /go/src/api

RUN go get github.com/go-sql-driver/mysql \
  && go get -u github.com/gin-gonic/gin
