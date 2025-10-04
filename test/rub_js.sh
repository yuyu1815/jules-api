#!/bin/bash

set -e

echo "Running JavaScript API tests..."

cd "$(dirname "$0")"

npm test

echo "JavaScript tests completed."
