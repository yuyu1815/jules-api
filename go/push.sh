#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

if [ $# -eq 0 ]; then
    echo -e "${RED}Error: Version not provided${NC}"
    echo "Usage: ./push.sh <version>"
    exit 1
fi

VERSION=$1

# Build and test
go build ./...
go test ./...

# Commit version changes if needed
# git add .
# git commit -m "Release v$VERSION"

# Create and push tag
echo -e "${YELLOW}Creating tag v$VERSION...${NC}"
git tag "v$VERSION"
git push origin "v$VERSION"

echo -e "${GREEN}Go module published successfully!${NC}"
echo "Users can now install with: go get github.com/yuyu1815/jules-api/go@v$VERSION"
