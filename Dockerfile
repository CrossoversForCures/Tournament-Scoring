# Start with a lightweight Go base image
FROM golang:1.22-alpine

# Set the working directory in the container
WORKDIR /app

# Copy dependency files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application will listen on
EXPOSE 8000

# Command to start the application
CMD ["./main"]