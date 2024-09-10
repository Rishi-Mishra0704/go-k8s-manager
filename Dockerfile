# Build stage
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files first to leverage caching of dependencies
COPY go.mod go.sum ./

# Download the necessary Go dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o go-k8s-manager .

# Final stage (smaller production image)
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary and entrypoint script from the builder stage
COPY --from=builder /app/go-k8s-manager .
COPY entrypoint.sh .

# Make the entrypoint script executable
RUN chmod +x entrypoint.sh

# Expose the application port
EXPOSE 8080

# Use the entrypoint script to run the app
ENTRYPOINT ["./entrypoint.sh"]
