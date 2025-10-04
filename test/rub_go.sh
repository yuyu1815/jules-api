#!/bin/bash

set -e

echo "Running Go API tests..."

# Navigate to the directory of the script itself (e.g., /app/test)
cd "$(dirname "$0")"

# Ensure Go dependencies are tidy
echo "==> Tidying Go modules..."
go mod tidy

# Run the test program
echo "==> Running tests..."
go run test_api_go.go

echo "==> Go tests completed."