FROM golang:1.15.12 as build

RUN apt-get update

WORKDIR /go/src
COPY ./src /go/src

RUN go build -o go \
    && go get github.com/oxequa/realize \
    && go get -u github.com/go-delve/delve/cmd/dlv \
    && go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv \
    && go get github.com/rubenv/sql-migrate/... \
    && go get github.com/rakyll/statik