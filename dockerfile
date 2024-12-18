# Use the official Golang image as the base image for building the app
FROM golang:1.22 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules file
COPY go.mod go.sum ./

# Download all dependencies (this will cache dependencies if unchanged)
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o pandemonium_api ./cmd/server

# Use a newer Debian image or Ubuntu image with GLIBC 2.34+ for the runtime
FROM ubuntu:latest 

ENV NEXTCLOUD_USERNAME=${NEXTCLOUD_USERNAME}
ENV NEXTCLOUD_PASSWORD=${NEXTCLOUD_PASSWORD}

# Install necessary libraries (glibc and others)
RUN apt-get update && apt-get install -y libc6
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y mosquitto mosquitto-clients && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Expose the port your app will be listening on
EXPOSE 1883 8080

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/pandemonium_api /app/

CMD ["sh", "-c", "mosquitto -c /etc/mosquitto/mosquitto.conf & /app/pandemonium_api"]



