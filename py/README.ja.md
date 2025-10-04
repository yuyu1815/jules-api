# Jules API Python クライアント

[![PyPI version](https://badge.fury.io/py/jules-api.svg)](https://pypi.org/project/jules-api/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Python Versions](https://img.shields.io/pypi/pyversions/jules-api)

[Jules API](https://developers.google.com/jules/api) 公式 Python クライアントライブラリ。

## インストール

```bash
pip install jules-api
```

またはソースからインストール：

```bash
pip install -e .
```

## 使い方

```python
from jules_api import JulesClient
from jules_api import CreateSessionRequest, SourceContext, GithubRepoContext, SendMessageRequest

# APIキーを使用してクライアントを初期化
client = JulesClient(api_key="YOUR_API_KEY_HERE")

# 利用可能なソースをリスト
sources_response = client.list_sources()
print("利用可能なソース:")
for source in sources_response.sources:
    print(f"- {source.id}: {source.name}")
    if source.github_repo:
        print(f"  GitHub: {source.github_repo.owner}/{source.github_repo.repo}")

if not sources_response.sources:
    print("ソースが見つかりません。まず Jules WebアプリでGitHubリポジトリを接続してください。")
    exit(1)

# 新しいセッションを作成
first_source = sources_response.sources[0]
session_request = CreateSessionRequest(
    prompt="「Hello from Jules!」を表示するシンプルなWebアプリを作成します",
    source_context=SourceContext(
        source=first_source.name,
        github_repo_context=GithubRepoContext(starting_branch="main")
    ),
    title="Hello World アプリセッション"
)

session = client.create_session(session_request)
print(f"セッション作成: {session.id}")
print(f"タイトル: {session.title}")
print(f"プロンプト: {session.prompt}")

# エージェントにメッセージを送信
message_request = SendMessageRequest(
    prompt="もっと魅力的になるようにスタイルを追加してください。"
)
client.send_message(session.id, message_request)
print("メッセージをエージェントに送信しました！")

# アクティビティをリスト（オプション）
activities_response = client.list_activities(session.id, page_size=10)
print(f"\n{len(activities_response.activities)} 件のアクティビティが見つかりました:")
for activity in activities_response.activities:
    content = activity.content or "コンテンツなし"
    if len(content) > 100:
        content = content[:100] + "..."
    print(f"- {activity.type}: {content}")
```

## 代替クライアント作成

便利なファンクションも使用できます：

```python
from jules_api import create_client

client = create_client("YOUR_API_KEY_HERE")
```

## API リファレンス

### JulesClient

#### コンストラクタ

```python
client = JulesClient(api_key="your_api_key", base_url="https://jules.googleapis.com/v1alpha")
```

**パラメータ:**
- `api_key`: Jules APIキー（必須）
- `base_url`: API ベースURL（オプション、デフォルトは "https://jules.googleapis.com/v1alpha"）

#### メソッド

##### `list_sources(next_page_token=None)`

Jules アカウントで接続された利用可能なソースをすべてリストアップします。

**パラメータ:**
- `next_page_token`: ページネーション用のトークン（オプション）

**戻り値:** `ListSourcesResponse` オブジェクト、ソース配列とオプションの next_page_token を含む

##### `create_session(request)`

新しいセッションを作成します。

**パラメータ:**
- `request`: `CreateSessionRequest` オブジェクト、セッションパラメータを含む

**戻り値:** 作成されたセッションを表す `Session` オブジェクト

##### `list_sessions(page_size=None, next_page_token=None)`

あなたのセッションをリストアップします。

**パラメータ:**
- `page_size`: 返すセッションの最大数（オプション）
- `next_page_token`: ページネーション用のトークン（オプション）

**戻り値:** `ListSessionsResponse` オブジェクト、セッション配列とオプションの next_page_token を含む

##### `approve_plan(session_id)`

セッションの最新のプランを承認（require_plan_approval が True の場合に使用）。

**パラメータ:**
- `session_id`: セッションID文字列（必須）

##### `list_activities(session_id, page_size=None, next_page_token=None)`

セッションのアクティビティをリストアップします。

**パラメータ:**
- `session_id`: セッションID文字列（必須）
- `page_size`: 返すアクティビティの最大数（オプション）
- `next_page_token`: ページネーション用のトークン（オプション）

**戻り値:** `ListActivitiesResponse` オブジェクト、アクティビティ配列とオプションの next_page_token を含む

##### `send_message(session_id, request)`

セッション内のエージェントにメッセージを送信します。

**パラメータ:**
- `session_id`: セッションID文字列（必須）
- `request`: メッセージを含む `SendMessageRequest` オブジェクト

##### `get_session(session_id)`

特定のセッションの詳細を取得します。

**パラメータ:**
- `session_id`: セッションID文字列（必須）

**戻り値:** `Session` オブジェクト

##### `get_source(source_id)`

特定のソースの詳細を取得します。

**パラメータ:**
- `source_id`: ソースID文字列（必須）

**戻り値:** `Source` オブジェクト

## モデル

### リクエスト/レスポンスモデル

- `CreateSessionRequest`: 新しいセッション作成に使用
- `SendMessageRequest`: メッセージ送信に使用
- `ListSourcesResponse`: ソースリスト取得時のレスポンス
- `ListSessionsResponse`: セッションリスト取得時のレスポンス
- `ListActivitiesResponse`: アクティビティリスト取得時のレスポンス
- `Session`: セッション情報
- `Source`: ソース情報
- `Activity`: アクティビティ情報

### データモデル

- `SourceContext`: セッションソースのコンテキスト
- `GithubRepoContext`: GitHub リポジトリの追加コンテキスト
- `GithubRepo`: GitHub リポジトリ情報

すべてのモデルで Pydantic を使用して検証と型ヒントを行っています。

## 認証

すべての API リクエストには、Jules Web アプリの Settings ページから取得した Jules API キーを使用して認証が必要です。クライアントは自動的にキーを `X-Goog-Api-Key` ヘッダーのすべてのリクエストに含めます。

**重要:** API キーを安全に保管してください。バージョン管理にコミットしたり、クライアントサイドコードで公開したりしないでください。

環境変数として API キーを設定：

```bash
export JULES_API_KEY=your_api_key_here
```

## エラー処理

クライアントは HTTP エラーに対して `requests.HTTPError` 例外を発生させます。API 呼び出しは常に try-except ブロックでラップしてください：

```python
from requests.exceptions import HTTPError

try:
    session = client.create_session(request)
    print("成功:", session)
except HTTPError as e:
    print(f"HTTP エラー: {e}")
    print(f"ステータスコード: {e.response.status_code}")
    print(f"レスポンス: {e.response.text}")
```

## 型ヒント

このライブラリはすべての場所でモダンな Python 型ヒントを使用しています。あなたの IDE は優れたオートコンプリートと型チェックサポートを提供するはずです。

## 貢献

貢献を歓迎します！詳細は[コントリビューションガイドライン](../CONTRIBUTING.md)をお読みください。

## ライセンス

このプロジェクトは MIT ライセンスの下でライセンスされています - 詳細は[LICENSE](../LICENSE) ファイルをご覧ください。

## サポート

サポートについては、[support@jules.ai](mailto:support@jules.ai) までお問い合わせいただくか、[当社のドキュメント](https://developers.google.com/jules/api)をご覧ください。
