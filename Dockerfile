# Use an official Golang runtime as a parent image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build

# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["./employee-crud-api"]