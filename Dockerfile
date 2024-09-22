# Use official golang image as the base image
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum are unchanged.
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM debian:bullseye-slim

# Set environment variables for the app
ENV PORT=3000
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=goshop
ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432

# Install necessary libraries (Postgres SQL client for example)
RUN apt-get update && apt-get install -y \
    ca-certificates \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory in the second stage
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 3000

# Run the binary
CMD ["./main"]
