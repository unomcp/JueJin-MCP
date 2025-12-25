#!/usr/bin/env bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$PROJECT_ROOT"

echo "🚀 Building for all platforms..."
echo ""

# Clean dist directory
rm -rf dist
mkdir -p dist

# Build for all platforms
PLATFORMS=(
  "linux-amd64"
  "linux-arm64"
  "darwin-amd64"
  "darwin-arm64"
  "windows-amd64"
)

for platform in "${PLATFORMS[@]}"; do
  echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  bash "$SCRIPT_DIR/build-$platform.sh"
  echo ""
done

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ All builds completed!"
echo ""
echo "📦 Generated archives:"
ls -lh dist/*.{tar.gz,zip} 2>/dev/null || true

