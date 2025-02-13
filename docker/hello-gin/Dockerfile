# Use the official Golang image for building
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o hello-gin

# Use a smaller image for the final container
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/hello-gin .

# Expose the app port
EXPOSE 8080

# Run the application
CMD ["./hello-gin"]
