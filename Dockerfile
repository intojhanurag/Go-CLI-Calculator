FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download


COPY . .

RUN go build -o calculator ./cmd



FROM debian:bookworm-slim

WORKDIR /app


COPY --from=builder /app/calculator .

ENTRYPOINT ["./calculator"]