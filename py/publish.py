#!/usr/bin/env python3
"""
Publish script for the Jules API Python client.

Usage:
  python publish.py [version]

- If a semantic version (e.g., 1.0.2) is provided, the script updates
  py/jules_api/__init__.py __version__ accordingly before building.
- Requires twine configured (via ~/.pypirc or TWINE_USERNAME/TWINE_PASSWORD).
"""

import subprocess
import sys
from pathlib import Path
import re


def run_command(cmd):
    """Run a shell command and check for errors."""
    print(f"Running: {' '.join(cmd)}")
    result = subprocess.run(cmd, check=True)
    return result


def set_version(new_version: str):
    here = Path(__file__).resolve().parent
    init_path = here / "jules_api" / "__init__.py"
    content = init_path.read_text(encoding="utf-8")
    content_new = re.sub(
        r"^(__version__\s*=\s*)['\"][^'\"]+['\"]",
        rf"\1'{new_version}'",
        content,
        flags=re.MULTILINE,
    )
    if content == content_new:
        raise RuntimeError("Failed to update version in jules_api/__init__.py")
    init_path.write_text(content_new, encoding="utf-8")
    print(f"Bumped version to {new_version}")


def main():
    """Build and publish the package."""
    try:
        # Optional version bump
        if len(sys.argv) > 1:
            version = sys.argv[1].strip()
            if not re.match(r"^[0-9]+\.[0-9]+\.[0-9]+[a-z0-9\.-]*$", version):
                raise ValueError(f"Invalid version string: {version}")
            set_version(version)

        # Clean previous builds
        run_command(["rm", "-rf", "dist/", "build/", "*.egg-info"])

        # Build the package
        run_command([sys.executable, "setup.py", "sdist", "bdist_wheel"])

        # Publish to PyPI (requires API token in ~/.pypirc or TWINE_USERNAME/TWINE_PASSWORD)
        run_command(["twine", "upload", "dist/*"])

        print("Package published successfully!")

    except subprocess.CalledProcessError as e:
        raise RuntimeError(f"Error during publishing: {e}")


if __name__ == "__main__":
    main()
