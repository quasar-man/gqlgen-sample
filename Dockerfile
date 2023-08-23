FROM golang:latest

WORKDIR /api
COPY ./ /api

RUN go mod download

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 8080

CMD ["go", "run", "server.go"]
