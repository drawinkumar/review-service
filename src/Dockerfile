# syntax=docker/dockerfile:1
FROM golang:1.23
WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN go mod download

RUN go build -o /godocker
WORKDIR /app

EXPOSE 8080

# wait 10 seconds for mysql to start
CMD echo "\n\n\nwaiting 10 seconds for mysql and minio to start...\n\n\n" && sleep 10 && /godocker