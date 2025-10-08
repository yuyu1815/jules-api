# Jules API Python 客户端

[![PyPI version](https://badge.fury.io/py/jules-api.svg)](https://pypi.org/project/jules-api/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Python Versions](https://img.shields.io/pypi/pyversions/jules-api)

道歉：我们在早期版本中误将其呈现为官方库。我们对此表示歉意。这是非官方库。

Jules API 的非官方 Python 客户端库。

## 安装

```bash
pip install jules-api
```

## 使用

```python
from jules_api import JulesClient
from jules_api import CreateSessionRequest, SourceContext, GithubRepoContext, SendMessageRequest

# 使用 API 密钥初始化客户端
client = JulesClient(api_key="YOUR_API_KEY_HERE")

# 列出可用源码
sources_response = client.list_sources()
print("可用源码:")
for source in sources_response.sources:
    print(f"- {source.id}: {source.name}")
    if source.github_repo:
        print(f"  GitHub: {source.github_repo.owner}/{source.github_repo.repo}")

if not sources_response.sources:
    print("未找到源码。请首先在 Jules Web 应用中连接 GitHub 仓库。")
    exit(1)

# 创建新会话
first_source = sources_response.sources[0]
session_request = CreateSessionRequest(
    prompt="创建一个显示 'Hello from Jules!' 的简单 Web 应用",
    source_context=SourceContext(
        source=first_source.name,
        github_repo_context=GithubRepoContext(starting_branch="main")
    ),
    title="Hello World 应用会话"
)

session = client.create_session(session_request)
print(f"创建会话: {session.id}")
print(f"标题: {session.title}")
print(f"提示: {session.prompt}")

# 向代理发送消息
message_request = SendMessageRequest(
    prompt="请添加一些样式使其更具吸引力。"
)
client.send_message(session.id, message_request)
print("消息已发送给代理！")

# 列出活动（可选）
activities_response = client.list_activities(session.id, page_size=10)
print(f"\n找到 {len(activities_response.activities)} 个活动:")
for activity in activities_response.activities:
    content = activity.content or "无内容"
    if len(content) > 100:
        content = content[:100] + "..."
    print(f"- {activity.type}: {content}")
```

## 替代客户端创建

您也可以使用便捷函数：

```python
from jules_api import create_client

client = create_client("YOUR_API_KEY_HERE")
```

## API 参考

### JulesClient

#### 构造函数

```python
client = JulesClient(api_key="your_api_key", base_url="https://jules.googleapis.com/v1alpha")
```

**参数:**
- `api_key`: Jules API 密钥（必需）
- `base_url`: API 基础 URL（可选，默认值 "https://jules.googleapis.com/v1alpha"）

#### 方法

##### `list_sources(next_page_token=None)`

列出连接到您的 Jules 账户的所有可用源码。

**参数:**
- `next_page_token`: 分页令牌（可选）

**返回:** 包含源数组和可选 next_page_token 的 `ListSourcesResponse` 对象

##### `create_session(request)`

创建新会话。

**参数:**
- `request`: 包含会话参数的 `CreateSessionRequest` 对象

**返回:** 表示创建会话的 `Session` 对象

##### `list_sessions(page_size=None, next_page_token=None)`

列出您的会话。

**参数:**
- `page_size`: 要返回的最大会话数（可选）
- `next_page_token`: 分页令牌（可选）

**返回:** 包含会话数组和可选 next_page_token 的 `ListSessionsResponse` 对象

##### `approve_plan(session_id)`

批准会话的最新计划（当 require_plan_approval 为 True 时使用）。

**参数:**
- `session_id`: 会话 ID 字符串（必需）

##### `list_activities(session_id, page_size=None, next_page_token=None)`

列出会话活动。

**参数:**
- `session_id`: 会话 ID 字符串（必需）
- `page_size`: 要返回的最大活动数（可选）
- `next_page_token`: 分页令牌（可选）

**返回:** 包含活动数组和可选 next_page_token 的 `ListActivitiesResponse` 对象

##### `send_message(session_id, request)`

向会话中的代理发送消息。

**参数:**
- `session_id`: 会话 ID 字符串（必需）
- `request`: 包含消息的 `SendMessageRequest` 对象

##### `get_session(session_id)`

获取特定会话的详细信息。

**参数:**
- `session_id`: 会话 ID 字符串（必需）

**返回:** `Session` 对象

##### `get_source(source_id)`

获取特定源码的详细信息。

**参数:**
- `source_id`: 源 ID 字符串（必需）

**返回:** `Source` 对象

## 模型

### 请求/响应模型

- `CreateSessionRequest`: 用于创建新会话
- `SendMessageRequest`: 用于发送消息
- `ListSourcesResponse`: 列出源码时的响应
- `ListSessionsResponse`: 列出会话时的响应
- `ListActivitiesResponse`: 列出活动时的响应
- `Session`: 会话信息
- `Source`: 源码信息
- `Activity`: 活动信息

### 数据模型

- `SourceContext`: 会话源码的上下文
- `GithubRepoContext`: GitHub 仓库的附加上下文
- `GithubRepo`: GitHub 仓库信息

所有模型使用 Pydantic 进行验证和类型提示。

## 身份验证

所有 API 请求都需要使用从 Jules Web 应用 Settings 页面获取的 Jules API 密钥进行身份验证。客户端自动在所有请求的 `X-Goog-Api-Key` 头中包含密钥。

**重要:** 请安全保管您的 API 密钥。切勿将其提交到版本控制或在客户端代码中公开。

将您的 API 密钥设置为环境变量：

```bash
export JULES_API_KEY=your_api_key_here
```

## 错误处理

客户端为 HTTP 错误引发 `requests.HTTPError` 异常。请始终将 API 调用包装在 try-except 块中：

```python
from requests.exceptions import HTTPError

try:
    session = client.create_session(request)
    print("成功:", session)
except HTTPError as e:
    print(f"HTTP 错误: {e}")
    print(f"状态码: {e.response.status_code}")
    print(f"响应: {e.response.text}")
```

## 类型提示

此库在所有地方使用现代 Python 类型提示。您的 IDE 应该提供出色的自动完成功能和类型检查支持。

## 贡献

欢迎贡献！请阅读我们的[贡献准则](../CONTRIBUTING.md)了解详情。

## 许可证

本项目根据 MIT 许可证获得许可 - 详情请参阅[LICENSE](../LICENSE) 文件。
