# -------------------------------------------
# Build stage: build the Go binary
# -------------------------------------------
FROM golang:1.22 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first — for caching dependencies
COPY go.mod go.sum ./

# Download modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o ingestion-service main.go

# -------------------------------------------
# Final stage: a minimal runtime image
# -------------------------------------------
FROM gcr.io/distroless/base-debian12

# Set working directory in runtime image
WORKDIR /

# Copy the binary from builder stage
COPY --from=builder /app/ingestion-service .

# Copy any config file if you have it
COPY config.yaml .

# Expose the port your Gin server uses
EXPOSE 8080

# Command to run the binary
ENTRYPOINT ["/ingestion-service"]
