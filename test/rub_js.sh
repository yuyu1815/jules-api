#!/bin/bash

set -e

echo "Running JavaScript API tests..."

# Navigate to the directory of the script itself (e.g., /app/test)
cd "$(dirname "$0")"

# Go to the js directory to build the client
echo "==> Building JS client..."
cd ../js
npm install > /dev/null 2>&1 # Suppress output for cleaner logs
npm run build
echo "==> Build complete."

# Go back to the test directory to run the tests
cd ../test

echo "==> Running tests..."
npm install > /dev/null 2>&1 # Install test dependencies (like dotenv)
node test_api_js.js

echo "==> JavaScript tests completed."