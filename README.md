[![PyPI Downloads](https://static.pepy.tech/personalized-badge/jules-api?period=total&units=INTERNATIONAL_SYSTEM&left_color=BLACK&right_color=GREEN&left_text=downloads)](https://pepy.tech/projects/jules-api)

# Jules API Client Libraries

Apology: We apologize for mistakenly presenting this library as official in earlier versions. This is an unofficial library.

This is an unofficial library.

This repository contains client libraries for the Jules API in three languages: JavaScript/Node.js, Go, and Python.

## Installation

Install the client library for your preferred language:

- **JavaScript/Node.js**: `npm install @yuzumican/jules-api` ([npm package](https://www.npmjs.com/package/@yuzumican/jules-api))
- **Python**: `pip install jules-api` ([PyPI package](https://pypi.org/project/jules-api/1.0/))
- **Go**: `go get github.com/yuyu1815/jules-api/go@latest` ([Go module](https://github.com/yuyu1815/jules-api/tree/main/go))

## About the Jules API

The Jules API lets you programmatically access Jules's capabilities to automate and enhance your software development lifecycle. You can use the API to create custom workflows, automate tasks like bug fixing and code reviews, and embed Jules's intelligence directly into the tools you use every day.

**Note:** The Jules API is in an alpha release, which means it is experimental.

## Libraries

- [**JavaScript/Node.js**](https://github.com/yuyu1815/jules-api/tree/main/js) - npm package with TypeScript support
- [**Go**](https://github.com/yuyu1815/jules-api/tree/main/go) - Go module
- [**Python**](https://github.com/yuyu1815/jules-api/tree/main/py) - Python package

Each library provides a complete client implementation for interacting with the Jules API, including:

- Authentication with API keys
- Source management
- Session management
- Activity handling
- Message sending

## Getting Started

Choose your preferred language and follow the instructions in the respective library's README.

## Authentication

All requests require an API key obtained from the Jules web app Settings page. The key should be passed in the `X-Goog-Api-Key` header.

Keep your API keys secure and never expose them in public code.

## API Concepts

- **Source**: Input source for the agent (e.g., GitHub repository)
- **Session**: Continuous unit of work within a specific context
- **Activity**: Single unit of work within a Session

## Important

This API is experimental and subject to change. For production use, wait for stable releases.

## Documentation

- [Japanese (日本語)](./README.ja.md)
- [Chinese (中文)](./README.zh.md)
