# Use an official Go image as the build environment
FROM golang:1.23-alpine AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the Go binary
COPY . .
RUN go build -o auth-service .

# Use a lightweight image for the final deployment
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the compiled Go binary from the build stage
COPY --from=build /app/auth-service .

# Set the environment variable to specify the port the service will run on
ENV PORT 3010

# Expose port 3010 to be used by the service
EXPOSE 3010

# Command to run the service
CMD ["./auth-service"]
