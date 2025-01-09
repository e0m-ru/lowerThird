# Use the official Golang image as the base image
FROM golang:1.17-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# # Copy the Go Modules files
# COPY go.mod ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o title .

# Expose port 3456 to the outside world
EXPOSE 3456

# Command to run the executable
CMD ["./title"]