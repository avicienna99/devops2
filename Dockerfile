# Use the official Go image to build and run the application
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests and source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]