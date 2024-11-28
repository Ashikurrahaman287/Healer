# Use golang:1.20-alpine as the base image for building the Go application
FROM golang:1.20-alpine as builder

# Set the working directory to /app
WORKDIR /app

# Copy go mod and sum files first for dependency caching
COPY go.mod go.sum ./

# Run go mod tidy to download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o healer .

# Use a smaller base image for running the application
FROM alpine:latest

# Install any required dependencies (e.g., ca-certificates)
RUN apk --no-cache add ca-certificates

# Set the working directory to /root
WORKDIR /root

# Copy the built binary from the builder image
COPY --from=builder /app/healer .

# Expose any ports if necessary (e.g., if the bot communicates over HTTP)
EXPOSE 8080

# Set the default command to run when the container starts
CMD ["./healer"]
