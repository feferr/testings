# Step 1: Build the Go binary
FROM golang:1.23-alpine AS build

# Install dependencies
RUN apk update && apk add --no-cache git bash

WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Step 2: Create the final image for production
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built Go binary from the build stage
COPY --from=build /app/main .

# Expose a port if necessary (for API or web-based CLI tools)
# EXPOSE 8080

# Run the app
CMD ["/app/main"]
