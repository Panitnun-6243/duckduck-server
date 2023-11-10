# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.21.3 as builder

# Copy local code to the container image.
WORKDIR /app
COPY . ./

# Build the binary, with all necessary external dependencies.
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o /server cmd/server/main.go

# Use a minimal Docker image to run the service binary.
FROM debian:bullseye-slim

# Install CA certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /server /server

EXPOSE 5050
# Run the web service on container startup.
CMD ["/server"]


