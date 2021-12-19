FROM golang:1.17.3-buster AS builder

WORKDIR /app
COPY . /app
RUN make build

FROM golang:1.17.3-buster
COPY --from=builder /go/bin/mangar /go/bin/mangar
EXPOSE 8080
CMD ["/go/bin/mangar"]
