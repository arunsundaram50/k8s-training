# --- Build Stage ---
FROM golang:1.17 AS build-stage

# Set working directory
WORKDIR /app

# Copy go mod and sum files
# COPY go.mod go.sum ./
COPY go.mod ./

# Copy the source code into the container
COPY hello.go .

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o hello.bin .

# --- Final Stage ---
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from build-stage
COPY --from=build-stage /app/hello.bin .

# Command to run the binary
CMD ["./hello.bin"]
