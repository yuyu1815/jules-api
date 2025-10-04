#!/usr/bin/env python3
"""
Publish script for the Jules API Python client.
"""

import subprocess
import sys

def run_command(cmd):
    """Run a shell command and check for errors."""
    print(f"Running: {' '.join(cmd)}")
    result = subprocess.run(cmd, check=True)
    return result

def main():
    """Build and publish the package."""
    try:
        # Clean previous builds
        run_command(["rm", "-rf", "dist/", "build/", "*.egg-info"])

        # Build the package
        run_command([sys.executable, "setup.py", "sdist", "bdist_wheel"])

        # Publish to PyPI (requires API token in ~/.pypirc or TWINE_USERNAME/TWINE_PASSWORD)
        run_command(["twine", "upload", "dist/*"])

        print("Package published successfully!")

    except subprocess.CalledProcessError as e:
        print(f"Error during publishing: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
