FROM golang:latest

ADD ./.docker/bin/entrypoint.dev.sh /usr/local/bin
WORKDIR /go/src/app
RUN go install github.com/Altoros/gorm-goose/cmd/gorm-goose@latest
RUN go install github.com/cosmtrek/air@latest
