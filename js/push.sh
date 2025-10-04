#!/bin/bash

set -e

# Ensure we're in the js directory
if [ ! -f package.json ]; then
  cd js
fi

npm publish --access=public

