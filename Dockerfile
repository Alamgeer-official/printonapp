# Use the official Golang image with Alpine Linux as the base image
FROM golang:1.21.5-alpine3.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o printonapp ./cmd/main.go

# Start a new stage from scratch
FROM alpine:3.19

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the current working directory
COPY --from=builder /app/printonapp .

# Expose port 4000 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["./printonapp"]
