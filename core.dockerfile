# Use an official Go image as the base image
FROM golang:1.22.7-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install necessary packages for the build
RUN apk add --no-cache git bash

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies are cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Copy the .env file to the container
COPY .env .env

# Install an environment variable loader (if needed)
RUN go install github.com/joho/godotenv/cmd/godotenv@latest

# Build the Go app
RUN go build -o main

# Create a minimal runtime image
FROM alpine:latest

FROM docker:dind

# Set the Current Working Directory inside the container
WORKDIR /root/

# Install bash, go, docker in the runtime image
RUN apk add --no-cache bash go docker

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file to the runtime image's working directory
COPY --from=builder /app/.env /root/

#copy the plugin.sh file to the runtime image's working directory
COPY --from=builder /app/plugin.sh /root/

# Make sure the binary is executable
RUN chmod +x main plugin.sh

# Expose the gRPC port
EXPOSE 50051

# Command to run the executable
CMD ["./main"]
