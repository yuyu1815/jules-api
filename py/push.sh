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

# Ensure we're in the py directory
if [ ! -f setup.py ]; then
  cd py
fi

# Update version in setup.py
sed -i.bak "s/version=\"[0-9]*\.[0-9]*\.[0-9]*a*\",/version=\"$VERSION\",/" setup.py

# Run tests using test/rub_py.sh
# ../test/rub_py.sh

# Install wheel for bdist_wheel support
pip install wheel

# Clean previous builds
rm -rf dist build *.egg-info

# Build the package
python3 setup.py sdist bdist_wheel

# Commit version changes if needed
# git add .
# git commit -m "Release v$VERSION"

# Create and push tag
echo -e "${YELLOW}Creating tag v$VERSION...${NC}"
git tag "v$VERSION"
git push origin "v$VERSION"

# Upload to PyPI (assuming twine is configured)
twine upload dist/*

echo -e "${GREEN}Python package published successfully!${NC}"
echo "Users can now install with: pip install jules-api==$VERSION"
