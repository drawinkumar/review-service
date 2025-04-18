# syntax=docker/dockerfile:1
FROM golang:1.23
WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN go mod download

RUN go build -o /godocker
WORKDIR /app

EXPOSE 8080

CMD sleep 3 && /godocker