# Start from a base image that includes Go runtime
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

RUN go mod tidy

# Build the Go application
RUN go build -o myapp ./app

# Expose a port on which the application will listen
EXPOSE 8084

# Set the command to run the executable when the container starts
CMD ["./myapp"]






