FROM golang:alpine

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go build ./cmd/main.go

CMD ["/app/main", "--config", "/app/config/dev.yml"]