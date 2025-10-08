# Jules API JavaScript 客户端

[![npm version](https://badge.fury.io/js/%40jules-ai%2Fjules-api.svg)](https://badge.fury.io/js/%40jules-ai%2Fjules-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

道歉：我们在早期版本中误将其呈现为官方库。我们对此表示歉意。这是非官方库。

Jules API 的非官方 Node.js/TypeScript 客户端库。

## 安装

```bash
npm install @yuzumican/jules-api
```

## 使用

```typescript
import { JulesClient } from '@yuzumican/jules-api';

// 使用 API 密钥初始化客户端
const client = new JulesClient({
  apiKey: 'YOUR_API_KEY_HERE'
});

// 列出可用源码
const sources = await client.listSources();
console.log('可用源码:', sources.sources);

// 创建新会话
const session = await client.createSession({
  prompt: '创建一个 boba 应用！',
  sourceContext: {
    source: 'sources/github/yourusername/yourrepo',
    githubRepoContext: {
      startingBranch: 'main'
    }
  },
  title: 'Boba 应用会话'
});
console.log('会话创建:', session);

// 向代理发送消息
await client.sendMessage(session.id, {
  prompt: '可以让它更像柯基犬吗？'
});

// 列出会话中的活动
const activities = await client.listActivities(session.id, 30);
console.log('会话活动:', activities.activities);
```

## API 参考

### JulesClient

#### 构造函数

```typescript
const client = new JulesClient(options: JulesClientOptions)
```

选项:
- `apiKey`: Jules API 密钥（必需）
- `baseUrl`: API 基础 URL（可选，默认 'https://jules.googleapis.com/v1alpha'）

#### 方法

##### `listSources(nextPageToken?: string): Promise<ListSourcesResponse>`

列出连接到您的 Jules 账户的所有可用源码。

**参数:**
- `nextPageToken`: 分页令牌（可选）

**返回:** 包含源数组和可选 nextPageToken 的对象

##### `createSession(request: CreateSessionRequest): Promise<Session>`

创建新会话。

**参数:**
- `request`: 会话创建参数

**CreateSessionRequest:**
- `prompt`: 会话的初始提示（必需）
- `sourceContext`: 包含源码名称和 GitHub 仓库详细信息的源码上下文（必需）
- `title`: 会话标题（必需）
- `requirePlanApproval`: 是否需要明确计划批准（可选，默认为 false）

**返回:** 创建的会话对象

##### `listSessions(pageSize?: number, nextPageToken?: string): Promise<ListSessionsResponse>`

列出您的会话。

**参数:**
- `pageSize`: 要返回的最大会话数（可选）
- `nextPageToken`: 分页令牌（可选）

**返回:** 包含会话数组和可选 nextPageToken 的对象

##### `approvePlan(sessionId: string): Promise<void>`

批准会话的最新计划（当 requirePlanApproval 为 true 时使用）。

**参数:**
- `sessionId`: 会话 ID（必需）

##### `listActivities(sessionId: string, pageSize?: number, nextPageToken?: string): Promise<ListActivitiesResponse>`

列出会话活动。

**参数:**
- `sessionId`: 会话 ID（必需）
- `pageSize`: 要返回的最大活动数（可选）
- `nextPageToken`: 分页令牌（可选）

**返回:** 包含活动数组和可选 nextPageToken 的对象

##### `sendMessage(sessionId: string, request: SendMessageRequest): Promise<void>`

向会话中的代理发送消息。

**参数:**
- `sessionId`: 会话 ID（必需）
- `request`: 消息参数

**SendMessageRequest:**
- `prompt`: 要发送的消息（必需）

##### `getSession(sessionId: string): Promise<Session>`

获取特定会话的详细信息。

**参数:**
- `sessionId`: 会话 ID（必需）

**返回:** 会话对象

##### `getSource(sourceId: string): Promise<Source>`

获取特定源码的详细信息。

**参数:**
- `sourceId`: 源 ID（必需）

**返回:** 源码对象

## 身份验证

所有 API 请求都需要使用从 Jules Web 应用 Settings 页面获取的 Jules API 密钥进行身份验证。客户端自动在所有请求的 `X-Goog-Api-Key` 头中包含密钥。

**重要:** 请安全保管您的 API 密钥。切勿将其提交到版本控制或在客户端代码中公开。

## 错误处理

客户端为 HTTP 错误抛出异常。请始终将 API 调用包装在 try-catch 块中：

```typescript
try {
  const session = await client.createSession(request);
  console.log('成功:', session);
} catch (error) {
  console.error('错误:', error.message);
}
```

## TypeScript 支持

此库使用 TypeScript 编写并包含完整的类型定义。您将获得出色的 IntelliSense 支持和编译时类型检查。

## 贡献

欢迎贡献！请阅读我们的[贡献准则](../CONTRIBUTING.md)了解详情。

## 许可证

本项目根据 MIT 许可证获得许可 - 详情请参阅[LICENSE](../LICENSE) 文件。
