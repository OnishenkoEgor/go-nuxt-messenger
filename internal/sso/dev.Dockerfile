FROM golang:1.25.4

WORKDIR /app

COPY go.mod ./

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go mod tidy
RUN go mod vendor

COPY *.go ./

EXPOSE 8080
