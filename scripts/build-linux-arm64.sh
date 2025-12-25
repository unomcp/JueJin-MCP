#!/usr/bin/env bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$PROJECT_ROOT"

GOOS=linux
GOARCH=arm64
BINARY_NAME="juejin-mcp"
OUTPUT_DIR="dist"

echo "Building for $GOOS/$GOARCH..."

# Create output directory
mkdir -p "$OUTPUT_DIR"

# Build
CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH \
  go build -v -ldflags="-s -w" -o "$OUTPUT_DIR/$BINARY_NAME" .

# Create archive
cd "$OUTPUT_DIR"
ARCHIVE_NAME="juejin-mcp-$GOOS-$GOARCH.tar.gz"
tar czf "$ARCHIVE_NAME" "$BINARY_NAME"

echo "✅ Build completed!"
echo "📦 Archive: $OUTPUT_DIR/$ARCHIVE_NAME"
ls -lh "$ARCHIVE_NAME"

