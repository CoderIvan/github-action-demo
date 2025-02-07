# syntax=docker/dockerfile:1

# Step 1: build golang binary
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY main.go ./
RUN go build -o start

# Step 2: copy binary from step1
FROM alpine
WORKDIR /app
COPY --from=builder /app/start /app/start
CMD ["/app/start"]