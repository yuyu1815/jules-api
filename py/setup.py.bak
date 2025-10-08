"""
Setup script for the Jules API Python client.
"""

from setuptools import setup, find_packages
from pathlib import Path
import re

here = Path(__file__).resolve().parent
readme_path = (here.parent / "README.md") if not (here / "README.md").exists() else (here / "README.md")

with open(readme_path, "r", encoding="utf-8") as fh:
    long_description = fh.read()

# Single-source the version from the package __init__.py
init_path = here / "jules_api" / "__init__.py"
version_match = re.search(r"^__version__\s*=\s*['\"]([^'\"]+)['\"]", init_path.read_text(encoding="utf-8"), re.MULTILINE)
if not version_match:
    raise RuntimeError("Unable to find __version__ in jules_api/__init__.py")
version = version_match.group(1)

setup(
    name="jules-api",
    version=version,
    author="Jules AI",
    author_email="support@jules.ai",
    description="Official Python client for the Jules API",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/yuyu1815/jules-api",
    packages=find_packages(),
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "Topic :: Software Development :: Libraries",
        "Topic :: Software Development :: Libraries :: Python Modules",
    ],
    python_requires=">=3.8",
    install_requires=[
        "requests>=2.25.0",
        "pydantic>=2.0.0",
    ],
    extras_require={
        "dev": [
            "pytest>=7.0.0",
            "pytest-asyncio>=0.21.0",
            "black>=23.0.0",
            "flake8>=6.0.0",
            "mypy>=1.0.0",
            "isort>=5.10.0",
            "python-dotenv>=1.0.0",
        ],
    },
)
