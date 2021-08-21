FROM golang:1.16.7-alpine3.14
ENV ENV GOMODULE=on
WORKDIR /app
COPY . /app
RUN apk update && \
    apk add --no-cache git && \
    go mod download
EXPOSE 8080
