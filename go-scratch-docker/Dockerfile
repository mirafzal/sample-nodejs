# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# Enable Go modules and CGO disabled for static binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY . .

RUN go build -o main .

# Stage 2: Create tiny final image
FROM scratch

# Copy statically compiled binary from builder
COPY --from=builder /app/main /main

# Expose port and set binary as entrypoint
EXPOSE 8080
ENTRYPOINT ["/main"]
