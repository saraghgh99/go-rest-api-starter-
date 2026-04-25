# Use the official Golang image as the base
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your project code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your Gin router runs on (usually 8080)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]