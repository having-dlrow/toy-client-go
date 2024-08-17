# Use an official Golang runtime as a parent image
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o ./app .

# Run the executable
CMD ["./app"]