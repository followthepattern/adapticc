# base Golang image
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /build

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
WORKDIR /out

# Create sock directory for unix socket communication
RUN mkdir /sock

# Copy cerbos binary
COPY --from=ghcr.io/cerbos/cerbos:0.34.0 /cerbos .

# Copy the binary from the builder stage to the final image
COPY --from=builder /build/app .

# Create configs folder to configs
RUN mkdir ./configs

# Copy configs
COPY ./configs/config.yaml.tmpl ./configs/config.yaml

# Create policies folder to cerbos config and policies
RUN mkdir /policies

# Copy cerbos config
COPY ./policies /policies

# Expose port 8080 for the application
EXPOSE 8080

# Set the command to run the binary
CMD ["sh", "-c", "./cerbos server --config=/policies/.cerbos.yaml & ./app"]