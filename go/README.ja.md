# Jules API Go クライアント

[![Go Reference](https://pkg.go.dev/badge/github.com/yuyu1815/jules-api/go.svg)](https://pkg.go.dev/github.com/yuyu1815/jules-api/go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

謝罪：以前のバージョンで誤ってこのライブラリを公式として紹介してしまい、申し訳ありません。このライブラリは非公式です。

Jules API の非公式 Go クライアントライブラリ。

## インストール

ライブラリをインストールするには、次のコマンドを実行してください：

```bash
go get github.com/yuyu1815/jules-api/go@latest
```


## 使い方

```go
package main

import (
	"fmt"
	"log"
	"os"

	jules "github.com/yuyu1815/jules-api/go"
)

func main() {
	// APIキーを使用してクライアントを初期化
	client := jules.NewClient(jules.NewClientOptions(os.Getenv("JULES_API_KEY")))

	// 利用可能なソースをリスト
	sources, err := client.ListSources("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("利用可能なソース:")
	for _, source := range sources.Sources {
		fmt.Printf("- %s: %s\n", source.ID, source.Name)
		if source.GithubRepo != nil {
			fmt.Printf("  GitHub: %s/%s\n", source.GithubRepo.Owner, source.GithubRepo.Repo)
		}
	}

	if len(sources.Sources) == 0 {
		fmt.Println("ソースが見つかりません。まず Jules WebアプリでGitHubリポジトリを接続してください。")
		return
	}

	// 新しいセッションを作成
	firstSource := sources.Sources[0]
	session, err := client.CreateSession(&jules.CreateSessionRequest{
		Prompt: "「Hello from Jules!」を表示するシンプルなWebアプリを作成します",
		SourceContext: jules.SourceContext{
			Source: firstSource.Name,
			GithubRepoContext: &jules.GithubRepoContext{
				StartingBranch: "main",
			},
		},
		Title: "Hello World アプリセッション",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("セッション作成: %s\n", session.ID)
	fmt.Printf("タイトル: %s\n", session.Title)
	fmt.Printf("プロンプト: %s\n", session.Prompt)

	// エージェントにメッセージを送信
	err = client.SendMessage(session.ID, &jules.SendMessageRequest{
		Prompt: "もっと魅力的になるようにスタイルを追加してください。",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("メッセージをエージェントに送信しました！")
}
```

## API リファレンス

### Client

#### コンストラクタ

```go
func NewClient(options *ClientOptions) *Client
```

新しい Jules API クライアントを作成します。

**ClientOptions:**
- `APIKey`: Jules APIキー（必須）
- `BaseURL`: API ベースURL（オプション、デフォルトは "https://jules.googleapis.com/v1alpha"）

#### メソッド

##### `ListSources(nextPageToken string) (*ListSourcesResponse, error)`

Jules アカウントで接続された利用可能なソースをすべてリストアップします。

**パラメータ:**
- `nextPageToken`: ページネーション用のトークン（空文字列でも可）

**戻り値:** ソース配列とオプションの nextPageToken を含む ListSourcesResponse

##### `CreateSession(request *CreateSessionRequest) (*Session, error)`

新しいセッションを作成します。

**パラメータ:**
- `request`: セッション作成パラメータ

**CreateSessionRequest:**
- `Prompt`: セッションの初期プロンプト（必須）
- `SourceContext`: ソース名と GitHub リポジトリの詳細を含むソースコンテキスト（必須）
- `Title`: セッションタイトル（必須）
- `RequirePlanApproval`: 計画の明示的な承認を必要とするかどうか（オプション、デフォルトは false）

**戻り値:** 作成されたセッションオブジェクト

##### `ListSessions(pageSize int, nextPageToken string) (*ListSessionsResponse, error)`

あなたのセッションをリストアップします。

**パラメータ:**
- `pageSize`: 返すセッションの最大数（0でデフォルト）
- `nextPageToken`: ページネーション用のトークン（空文字列でも可）

**戻り値:** セッション配列とオプションの nextPageToken を含む ListSessionsResponse

##### `ApprovePlan(sessionID string) error`

セッションの最新のプランを承認（RequirePlanApproval が true の場合に使用）。

**パラメータ:**
- `sessionID`: セッションID（必須）

##### `ListActivities(sessionID string, pageSize int, nextPageToken string) (*ListActivitiesResponse, error)`

セッションのアクティビティをリストアップします。

**パラメータ:**
- `sessionID`: セッションID（必須）
- `pageSize`: 返すアクティビティの最大数（0でデフォルト）
- `nextPageToken`: ページネーション用のトークン（空文字列でも可）

**戻り値:** アクティビティ配列とオプションの nextPageToken を含む ListActivitiesResponse

##### `SendMessage(sessionID string, request *SendMessageRequest) error`

セッション内のエージェントにメッセージを送信します。

**パラメータ:**
- `sessionID`: セッションID（必須）
- `request`: メッセージパラメータ

**SendMessageRequest:**
- `Prompt`: 送信するメッセージ（必須）

##### `GetSession(sessionID string) (*Session, error)`

特定のセッションの詳細を取得します。

**パラメータ:**
- `sessionID`: セッションID（必須）

**戻り値:** セッションオブジェクト

##### `GetSource(sourceID string) (*Source, error)`

特定のソースの詳細を取得します。

**パラメータ:**
- `sourceID`: ソースID（必須）

**戻り値:** ソースオブジェクト

## 認証

すべての API リクエストには、Jules Web アプリの Settings ページから取得した Jules API キーを使用して認証が必要です。クライアントは自動的にキーを `X-Goog-Api-Key` ヘッダーのすべてのリクエストに含めます。

**重要:** APIキーを安全に保管してください。バージョン管理にコミットしたり、パブリックコードで公開したりしないでください。

環境変数として API キーを設定：

```bash
export JULES_API_KEY=your_api_key_here
```

## エラー処理

クライアントは失敗した API リクエストに対して Go エラーを返します。エラーを常に確認してください：

```go
session, err := client.CreateSession(request)
if err != nil {
	log.Fatal(err)
}
```

## 貢献

貢献を歓迎します！詳細は[コントリビューションガイドライン](../CONTRIBUTING.md)をお読みください。

## ライセンス

このプロジェクトは MIT ライセンスの下でライセンスされています - 詳細は[LICENSE](../LICENSE) ファイルをご覧ください。
