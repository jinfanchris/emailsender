# Start from a minimal Go base image
FROM golang:1.24.2-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum to enable caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the full project
COPY . .

# Build the server binary
RUN go build -o /app/bin/server ./cmd/server
RUN go build -o /app/bin/client ./cmd/client
RUN go build -o /app/bin/example ./cmd/example

# Final minimal image
FROM alpine:latest

WORKDIR /root/

# Copy runtime files if needed
COPY runtime /root/runtime
COPY --from=builder /app/bin/server .
COPY --from=builder /app/bin/client .
COPY --from=builder /app/bin/example .

EXPOSE 52333

CMD ["./server"]
