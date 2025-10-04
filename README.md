# Jules API Client Libraries

This repository contains official client libraries for the Jules API in three languages: JavaScript/Node.js, Go, and Python.

## About the Jules API

The Jules API lets you programmatically access Jules's capabilities to automate and enhance your software development lifecycle. You can use the API to create custom workflows, automate tasks like bug fixing and code reviews, and embed Jules's intelligence directly into the tools you use every day.

**Note:** The Jules API is in an alpha release, which means it is experimental.

## Libraries

- [**JavaScript/Node.js**](./js) - npm package with TypeScript support
- [**Go**](./go) - Go module
- [**Python**](./py) - Python package

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
