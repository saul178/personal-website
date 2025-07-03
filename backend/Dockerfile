# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/api

# Stage 2: Runtime
FROM alpine:latest

# Optional: Add a non-root user (recommended for security)
RUN adduser -D -g '' appuser

WORKDIR /app

# Copy compiled binary
COPY --from=builder /app/server .

# Copy only the required runtime assets
COPY --from=builder /app/internal/assets /app/internal/assets
COPY --from=builder /app/internal/metadata /app/internal/metadata

EXPOSE 8080

# Switch to non-root user and give permission to app
RUN chown -R appuser:appuser /app
USER appuser

CMD ["./server"]
