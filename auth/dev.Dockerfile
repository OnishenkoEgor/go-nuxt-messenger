FROM golang:1.25.4

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

EXPOSE 8080
