FROM golang:1.21-alpine AS builder

# Install gcc and musl-dev because go-sqlite3 requires CGO
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build with CGO enabled (required for sqlite)
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o backend-app main.go

# Production stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/backend-app .
# Copy schema file
COPY --from=builder /app/database/schema.sql ./database/
# Copy assets directory (placeholder card, etc)
COPY --from=builder /app/assets ./assets

# Ensure data and uploads directories exist
RUN mkdir -p /app/data /app/uploads

EXPOSE 8080

CMD ["./backend-app"]
