# Ref: https://hub.docker.com/_/golang

FROM golang:1.24.5

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
