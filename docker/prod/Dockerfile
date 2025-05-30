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

# setup python dependencies
RUN pip3 install --break-system-packages \
    grpcio-tools

# setup typescript dependencies
RUN npm install \
    --save-dev @protobuf-ts/plugin \
    --save-dev @protobuf-ts/runtime
