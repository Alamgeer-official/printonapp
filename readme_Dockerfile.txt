# Use the official Golang image with Alpine Linux as the base image
FROM golang:1.21.5-alpine3.19

# Set the working directory inside the container
WORKDIR /printonapp/backend

# Copy the current directory contents into the container at /printonapp/backend
COPY . /printonapp/backend

# Build the Go application
RUN go build /printonapp/backend

# Expose port 4000 to the outside world
EXPOSE 4000

# Command to run the executable
ENTRYPOINT [ "./printonapp" ]
