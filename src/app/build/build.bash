#!/bin/bash

set -e 

OUTPUT_DIR=/src/app/build
mkdir -p "$OUTPUT_DIR"

platforms=("linux/amd64" "linux/arm64")

APP_NAME="protobuf-generator"

for platform in "${platforms[@]}"; do
  GOOS=${platform%/*}
  GOARCH=${platform#*/}
  output_name="${APP_NAME}-${GOOS}-${GOARCH}"

  env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -o "$OUTPUT_DIR/$output_name"
done

