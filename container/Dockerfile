# Build stage
FROM golang:latest AS build

WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy only the necessary artifacts from the build stage
COPY --from=build /app/main .

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
