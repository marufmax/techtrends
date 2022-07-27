FROM golang:1.18 as base

FROM base as dev

RUN apt install make

WORKDIR /data/collector

RUN go get && go build && go mod download

CMD ["go", "build", "cmd/main.go"]