# Use the official Golang image as a build stage
FROM golang:1.22-alpine as build

# Set the Current Working Directory inside the container
WORKDIR /simple-microservice-backend

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Copy the local.env file to the config directory inside the container
COPY config/local.env /simple-microservice-backend/config/local.env

# Set the working directory to the cmd folder where main.go is located
WORKDIR /simple-microservice-backend/cmd

# Build the Go app
RUN go build -o /simple-microservice-backend/main .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build /simple-microservice-backend/main .

# Copy the config directory to the final stage (if needed by your app)
COPY --from=build /simple-microservice-backend/config /simple-microservice-backend/config

# Expose port 5200 to the outside world
EXPOSE 5200

# Command to run the executable
CMD ["./main"]
