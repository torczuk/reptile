FROM golang:1.14.1 AS builder
WORKDIR /go/src/github.com/torczuk/reptile/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o reptile .

FROM alpine:3.11.5
WORKDIR /reptile
COPY --from=builder /go/src/github.com/torczuk/reptile .
EXPOSE 2600
CMD ["./reptile"]
