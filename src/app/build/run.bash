#!/bin/bash

set -e

ARCH=$(uname -m)

case "$ARCH" in
  x86_64)
    BINARY="/src/app/build/protobuf-generator-linux-amd64"
    ;;
  aarch64 | arm64)
    BINARY="/src/app/build/protobuf-generator-linux-arm64"
    ;;
  *)
    echo "[ERROR] not supported architecture: $ARCH"
    exit 1
    ;;
esac

exec "$BINARY" "$@"
