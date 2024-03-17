FROM golang:1.21

RUN apt-get update && apt-get install -y \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir /sock

RUN mkdir /sidecar

RUN curl -L -o /sidecar/cerbos.tar.gz "https://github.com/cerbos/cerbos/releases/download/v0.34.0/cerbos_0.34.0_Linux_arm64.tar.gz"

RUN tar -xzvf /sidecar/cerbos.tar.gz -C /sidecar

RUN chmod +x /sidecar/cerbos

VOLUME ["/policies", "/tmp", "/.cache"]