FROM golang:1.21.6

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy everything from the current directory to the working directory inside the container
COPY . .

# Download and install any dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
