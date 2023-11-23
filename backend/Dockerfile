# Stage 1: Build Golang binary
FROM golang:1.20-alpine AS build

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Golang application
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Stage 2: Create a minimal image to run the Golang app
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the previous stage
COPY --from=build /app/app .

# Expose port used by the Golang application
EXPOSE 8080

# Command to run the Golang application
CMD ["./app"]
