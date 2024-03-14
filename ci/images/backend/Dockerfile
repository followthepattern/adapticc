# base Golang image
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /dev

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files to the container
COPY . .

# Build the application with CGO_ENABLED=0 to create a static binary
RUN CGO_ENABLED=0 go build -o app ./cmd/adapticc

# Start with a minimal image to reduce the attack surface
FROM alpine:3.19.1

# Set the working directory
WORKDIR /prod

# Copy the binary from the builder stage to the final image
COPY --from=builder /dev/app .

# Create configs folder to configs
RUN mkdir ./configs

# Copy configs from builder
COPY --from=builder /dev/configs/config.yaml.tmpl ./configs/config.yaml

# Expose port 8080 for the application
EXPOSE 8080

# Set the command to run the binary
CMD ["./app"]