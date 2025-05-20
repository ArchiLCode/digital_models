FROM golang:1.21-alpine

WORKDIR /app

# Install required system packages
RUN apk add --no-cache postgresql-client

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code and migrations
COPY . .

# Ensure migrations directory exists
RUN mkdir -p migrations

# Build the application
RUN go build -o main cmd/main.go

# Command to run the application
CMD ["./main"]