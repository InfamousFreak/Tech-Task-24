# Build stage
FROM golang:alpine AS builder

# Install Git
RUN apk add --no-cache git

# Set the working directory
WORKDIR /go/src/app

# Copy the source code into the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o /go/bin/app .

# Final stage
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Copy the built Go app from the build stage
COPY --from=builder /go/bin/app .bin//app

# Copy the .env file
COPY .env .env

# Set the entry point to run the app
ENTRYPOINT [".bin/app"]

# Label the image
LABEL Name=techtask24 Version=0.0.1

# Expose the port the app runs on
EXPOSE 8080
