# Use the official Golang image as the base image
FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy the go mod and sum files first for caching dependencies
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the entire application source code to the container
COPY . .

# Build the Go application
RUN go build -o main main.go

# Command to run the executable
CMD ["./main"]
