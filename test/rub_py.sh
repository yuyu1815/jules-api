#!/bin/bash

set -e

echo "Running Python API tests..."

cd "$(dirname "$0")"

# Create virtual environment if it doesn't exist
if [ ! -d ".venv" ]; then
    echo "Creating virtual environment..."
    python3 -m venv .venv
fi

# Activate virtual environment
echo "Activating virtual environment..."
source .venv/bin/activate

# Install dependencies
echo "Installing dependencies..."
pip install -r requirements.txt

# Run the test
echo "Running tests..."
python3 test_api.py

echo "Python tests completed."
