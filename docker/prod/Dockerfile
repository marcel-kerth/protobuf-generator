FROM debian:bookworm-slim

WORKDIR /root

# setup os dependencies
RUN apt update && apt install -y \
    curl \
    tar \
    nodejs \
    npm \
    python3 \
    python3-pip \
    protobuf-compiler \
    protoc-gen-go \
    protoc-gen-go-grpc

# install go | ARM64
ENV GO_VERSION=1.24.3
ENV PATH="/usr/local/go/bin:/root/go/bin:$PATH"
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-arm64.tar.gz -o go.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

# setup python dependencies
RUN pip3 install --break-system-packages \
    grpcio-tools

# setup typescript dependencies
RUN npm install \
    --save-dev @protobuf-ts/plugin \
    --save-dev @protobuf-ts/runtime
