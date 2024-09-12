# Use the official Go image as the base image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY *.go ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

# Use a minimal alpine image for the final stage
FROM alpine:3.14

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["/app/main"]
