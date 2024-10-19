# Step 1: Build stage
# Use an official Go image as the build environment
FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server -ldflags="-s -w" .

# Step 2: Runtime stage
# Use a smaller, lightweight base image like alpine for the final container
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/index.html .

# Expose the application's port
EXPOSE 8000

# Command to run the executable
CMD ["./server"]


