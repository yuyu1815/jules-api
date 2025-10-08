# Jules API Go 客户端

[![Go Reference](https://pkg.go.dev/badge/github.com/yuyu1815/jules-api/go.svg)](https://pkg.go.dev/github.com/yuyu1815/jules-api/go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

道歉：我们在早期版本中误将其呈现为官方库。我们对此表示歉意。这是非官方库。

Jules API 的非官方 Go 客户端库。

## 安装

运行以下命令安装库：

```bash
go get github.com/yuyu1815/jules-api/go@latest
```

## 使用

```go
package main

import (
	"fmt"
	"log"
	"os"

	jules "github.com/yuyu1815/jules-api/go"
)

func main() {
	// 使用 API 密钥初始化客户端
	client := jules.NewClient(jules.NewClientOptions(os.Getenv("JULES_API_KEY")))

	// 列出可用源码
	sources, err := client.ListSources("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("可用源码:")
	for _, source := range sources.Sources {
		fmt.Printf("- %s: %s\n", source.ID, source.Name)
		if source.GithubRepo != nil {
			fmt.Printf("  GitHub: %s/%s\n", source.GithubRepo.Owner, source.GithubRepo.Repo)
		}
	}

	if len(sources.Sources) == 0 {
		fmt.Println("未找到源码。请首先在 Jules Web 应用中连接 GitHub 仓库。")
		return
	}

	// 创建新会话
	firstSource := sources.Sources[0]
	session, err := client.CreateSession(&jules.CreateSessionRequest{
		Prompt: "创建一个显示 'Hello from Jules!' 的简单 Web 应用",
		SourceContext: jules.SourceContext{
			Source: firstSource.Name,
			GithubRepoContext: &jules.GithubRepoContext{
				StartingBranch: "main",
			},
		},
		Title: "Hello World 应用会话",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("创建会话: %s\n", session.ID)
	fmt.Printf("标题: %s\n", session.Title)
	fmt.Printf("提示: %s\n", session.Prompt)

	// 向代理发送消息
	err = client.SendMessage(session.ID, &jules.SendMessageRequest{
		Prompt: "请添加一些样式使其更具吸引力。",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("消息已发送给代理！")
}
```

## API 参考

### Client

#### 构造函数

```go
func NewClient(options *ClientOptions) *Client
```

创建新的 Jules API 客户端。

**ClientOptions:**
- `APIKey`: Jules API 密钥（必需）
- `BaseURL`: API 基础 URL（可选，默认 "https://jules.googleapis.com/v1alpha"）

#### 方法

##### `ListSources(nextPageToken string) (*ListSourcesResponse, error)`

列出连接到您的 Jules 账户的所有可用源码。

**参数:**
- `nextPageToken`: 分页令牌（可以为空字符串）

**返回:** 包含源数组和可选 nextPageToken 的 ListSourcesResponse

##### `CreateSession(request *CreateSessionRequest) (*Session, error)`

创建新会话。

**参数:**
- `request`: 会话创建参数

**CreateSessionRequest:**
- `Prompt`: 会话的初始提示（必需）
- `SourceContext`: 包含源码名称和 GitHub 仓库详细信息的源码上下文（必需）
- `Title`: 会话标题（必需）
- `RequirePlanApproval`: 是否需要明确计划批准（可选，默认为 false）

**返回:** 创建的会话对象

##### `ListSessions(pageSize int, nextPageToken string) (*ListSessionsResponse, error)`

列出您的会话。

**参数:**
- `pageSize`: 要返回的最大会话数（0 为默认）
- `nextPageToken`: 分页令牌（可以为空字符串）

**返回:** 包含会话数组和可选 nextPageToken 的 ListSessionsResponse

##### `ApprovePlan(sessionID string) error`

批准会话的最新计划（当 RequirePlanApproval 为 true 时使用）。

**参数:**
- `sessionID`: 会话 ID（必需）

##### `ListActivities(sessionID string, pageSize int, nextPageToken string) (*ListActivitiesResponse, error)`

列出会话活动。

**参数:**
- `sessionID`: 会话 ID（必需）
- `pageSize`: 要返回的最大活动数（0 为默认）
- `nextPageToken`: 分页令牌（可以为空字符串）

**返回:** 包含活动数组和可选 nextPageToken 的 ListActivitiesResponse

##### `SendMessage(sessionID string, request *SendMessageRequest) error`

向会话中的代理发送消息。

**参数:**
- `sessionID`: 会话 ID（必需）
- `request`: 消息参数

**SendMessageRequest:**
- `Prompt`: 要发送的消息（必需）

##### `GetSession(sessionID string) (*Session, error)`

获取特定会话的详细信息。

**参数:**
- `sessionID`: 会话 ID（必需）

**返回:** 会话对象

##### `GetSource(sourceID string) (*Source, error)`

获取特定源码的详细信息。

**参数:**
- `sourceID`: 源 ID（必需）

**返回:** 源码对象

## 身份验证

所有 API 请求都需要使用从 Jules Web 应用 Settings 页面获取的 Jules API 密钥进行身份验证。客户端自动在所有请求的 `X-Goog-Api-Key` 头中包含密钥。

**重要:** 请安全保管您的 API 密钥。切勿将其提交到版本控制或在公共代码中公开。

将您的 API 密钥设置为环境变量：

```bash
export JULES_API_KEY=your_api_key_here
```

## 错误处理

客户端为失败的 API 请求返回 Go 错误。请始终检查错误：

```go
session, err := client.CreateSession(request)
if err != nil {
	log.Fatal(err)
}
```

## 贡献

欢迎贡献！请阅读我们的[贡献准则](../CONTRIBUTING.md)了解详情。

## 许可证

本项目根据 MIT 许可证获得许可 - 详情请参阅[LICENSE](../LICENSE) 文件。
