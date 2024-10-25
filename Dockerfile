# Build stage
FROM golang:1.23.2-alpine AS builder

# Install git and make
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the main application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Build the database CLI tool
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o db-cli cmd/db/main.go

# Final stage
FROM alpine:3.18

# Add necessary runtime packages
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy the binaries from builder
COPY --from=builder /app/main .
COPY --from=builder /app/db-cli .
COPY --from=builder /app/.env .

# Create a non-root user
RUN adduser -D -g '' appuser
RUN chown -R appuser:appuser /app
USER appuser

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]
