FROM golang:1.22.0-bullseye

RUN apt-get update \
    && apt-get install -y git
RUN go install github.com/cosmtrek/air@latest

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2

WORKDIR /app

CMD ["air", "-c", ".air.toml"]
