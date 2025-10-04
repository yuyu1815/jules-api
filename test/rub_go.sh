#!/bin/bash

set -e

echo "Running Go API tests..."

cd "$(dirname "$0")"

go run test_api_go.go

echo "Go tests completed."
