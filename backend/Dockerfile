FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build with CGO disabled (glebarez/go-sqlite is pure Go!)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o backend-app main.go

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
# Copy wilayah data (pre-fetched, no HTTP needed at startup)
COPY --from=builder /app/data/wilayah.json ./data/wilayah.json

EXPOSE 8080

CMD ["./backend-app"]
