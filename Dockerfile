# Use the official Golang image as the base image
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o coubcore .

# Use a minimal Alpine image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a user group and user
RUN addgroup -g 1001 -S coubcore &&\
    adduser -u 1001 -S coubcore -G coubcore

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/coubcore .

# Copy the blockchain database if it exists
COPY --from=builder /app/blockchain.db . || echo "No blockchain.db to copy"

# Change ownership of the files to the coubcore user
RUN chown -R coubcore:coubcore /app

# Switch to the coubcore user
USER coubcore

# Expose the port that the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./coubcore"]