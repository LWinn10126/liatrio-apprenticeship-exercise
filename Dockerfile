### Stage 1: Build the Go application ###
FROM golang:1.26 AS builder

# Set destination for COPY 
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the Source code.  
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

### Stage 2: Create a minimal image to run the application ###
FROM scratch

# Copy the built Go application from the builder stage
COPY --from=builder /app/server /server

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./server"]