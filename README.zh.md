# Jules API 客户端库

道歉：我们在早期版本中误将其呈现为官方库。我们对此表示歉意。这是非官方库。

这是非官方库。

此仓库包含用于 Jules API 的客户端库，支持三种语言：JavaScript/Node.js、Go 和 Python。

## 安装

安装适合您首选语言的客户端库：

- **JavaScript/Node.js**: `npm install @yuzumican/jules-api` ([npm 软件包](https://www.npmjs.com/package/@yuzumican/jules-api))
- **Python**: `pip install jules-api` ([PyPI 软件包](https://pypi.org/project/jules-api/1.0/))
- **Go**: `go get github.com/yuyu1815/jules-api/go@latest` ([Go 模块](https://github.com/yuyu1815/jules-api/tree/main/go))

## 关于 Jules API

Jules API 允许您以编程方式访问 Jules 的功能，以自动化和增强您的软件开发生命周期。您可以使用 API 创建自定义工作流程，自动化任务，如错误修复和代码审查，并将 Jules 的智能直接嵌入到您日常使用的工具中。

**注意：** Jules API 处于 alpha 发布阶段，这意味着它是实验性的。

## 库

- [**JavaScript/Node.js**](https://github.com/yuyu1815/jules-api/tree/main/js) - npm 包，支持 TypeScript
- [**Go**](https://github.com/yuyu1815/jules-api/tree/main/go) - Go 模块
- [**Python**](https://github.com/yuyu1815/jules-api/tree/main/py) - Python 包

每个库提供 Jules API 交互的完整客户端实现，包括：

- 使用 API 密钥进行身份验证
- 源管理
- 会话管理
- 活动处理
- 消息发送

## 入门

选择您的首选语言并按照相应库的 README 中的说明进行操作。

## 身份验证

所有请求都需要通过 Jules Web 应用 Settings 页面获取的 API 密钥。在 `X-Goog-Api-Key` 头中传递密钥。

将您的 API 密钥保密，不要在公共代码中公开。

## API 概念

- **源**：代理的输入源（例如，GitHub 仓库）
- **会话**：特定上下文中的连续工作单元
- **活动**：会话中的单个工作单元

## 重要

此 API 是实验性的，可能会发生变化。对于生产使用，请等待稳定版本的发布。

## 文档

- [Japanese (日本語)](./README.ja.md)
- [Chinese (中文)](./README.zh.md)
