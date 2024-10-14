FROM golang:alpine

# Install git
RUN apk update && apk add --no-cache git

# Setup working directory
WORKDIR /app

# Copy the source code
COPY . .
COPY .env .

# Download and install the dependencies
RUN go get -d -v ./...
RUN go install -v ./...

# Build the Go application
RUN go build -o /vhtask ./cmd/main.go

# Expose the port
EXPOSE 8080

# Run the application
CMD ["/vhtask"]
