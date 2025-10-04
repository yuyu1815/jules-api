# Jules API JavaScript クライアント

[![npm version](https://badge.fury.io/js/%40jules-ai%2Fjules-api.svg)](https://badge.fury.io/js/%40jules-ai%2Fjules-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Jules API の非公式 Node.js/TypeScript クライアントライブラリ。

## インストール

```bash
npm install @yuzumican/jules-api
```

## 使い方

```typescript
import { JulesClient } from '@yuzumican/jules-api';

// APIキーを使用してクライアントを初期化
const client = new JulesClient({
  apiKey: 'YOUR_API_KEY_HERE'
});

// 利用可能なソースをリスト
const sources = await client.listSources();
console.log('利用可能なソース:', sources.sources);

// 新しいセッションを作成
const session = await client.createSession({
  prompt: 'ボバアプリを作成！',
  sourceContext: {
    source: 'sources/github/yourusername/yourrepo',
    githubRepoContext: {
      startingBranch: 'main'
    }
  },
  title: 'ボバアプリセッション'
});
console.log('セッション作成:', session);

// エージェントにメッセージを送信
await client.sendMessage(session.id, {
  prompt: 'コルギー風にしてみませんか？'
});

// セッション内のアクティビティをリスト
const activities = await client.listActivities(session.id, 30);
console.log('セッションアクティビティ:', activities.activities);
```

## API リファレンス

### JulesClient

#### コンストラクタ

```typescript
const client = new JulesClient(options: JulesClientOptions)
```

オプション:
- `apiKey`: Jules APIキー（必須）
- `baseUrl`: API ベースURL（オプション、デフォルトは 'https://jules.googleapis.com/v1alpha'）

#### メソッド

##### `listSources(nextPageToken?: string): Promise<ListSourcesResponse>`

Jules アカウントで接続された利用可能なソースをすべてリストアップします。

**パラメータ:**
- `nextPageToken`: ページネーション用のトークン（オプション）

**戻り値:** ソース配列とオプションの nextPageToken を含むオブジェクト

##### `createSession(request: CreateSessionRequest): Promise<Session>`

新しいセッションを作成します。

**パラメータ:**
- `request`: セッション作成パラメータ

**CreateSessionRequest:**
- `prompt`: セッションの初期プロンプト（必須）
- `sourceContext`: ソース名とGitHubリポジトリの詳細を含むソースコンテキスト（必須）
- `title`: セッションタイトル（必須）
- `requirePlanApproval`: 計画の明示的な承認を必要とするかどうか（オプション、デフォルトは false）

**戻り値:** 作成されたセッションオブジェクト

##### `listSessions(pageSize?: number, nextPageToken?: string): Promise<ListSessionsResponse>`

あなたのセッションをリストアップします。

**パラメータ:**
- `pageSize`: 返すセッションの最大数（オプション）
- `nextPageToken`: ページネーション用のトークン（オプション）

**戻り値:** セッション配列とオプションの nextPageToken を含むオブジェクト

##### `approvePlan(sessionId: string): Promise<void>`

セッションの最新のプランを承認（requirePlanApproval が true の場合に使用）。

**パラメータ:**
- `sessionId`: セッションID（必須）

##### `listActivities(sessionId: string, pageSize?: number, nextPageToken?: string): Promise<ListActivitiesResponse>`

セッションのアクティビティをリストアップします。

**パラメータ:**
- `sessionId`: セッションID（必須）
- `pageSize`: 返すアクティビティの最大数（オプション）
- `nextPageToken`: ページネーション用のトークン（オプション）

**戻り値:** アクティビティ配列とオプションの nextPageToken を含むオブジェクト

##### `sendMessage(sessionId: string, request: SendMessageRequest): Promise<void>`

セッション内のエージェントにメッセージを送信します。

**パラメータ:**
- `sessionId`: セッションID（必須）
- `request`: メッセージパラメータ

**SendMessageRequest:**
- `prompt`: 送信するメッセージ（必須）

##### `getSession(sessionId: string): Promise<Session>`

特定のセッションの詳細を取得します。

**パラメータ:**
- `sessionId`: セッションID（必須）

**戻り値:** セッションオブジェクト

##### `getSource(sourceId: string): Promise<Source>`

特定のソースの詳細を取得します。

**パラメータ:**
- `sourceId`: ソースID（必須）

**戻り値:** ソースオブジェクト

## 認証

すべての API リクエストには、Jules Web アプリの Settings ページから取得した Jules API キーを使用して認証が必要です。クライアントは自動的にキーを `X-Goog-Api-Key` ヘッダーのすべてのリクエストに含めます。

**重要:** APIキーを安全に保管してください。バージョン管理にコミットしたり、クライアントサイドコードで公開したりしないでください。

## エラー処理

クライアントは HTTP エラーに対して例外をスローします。API 呼び出しは常に try-catch ブロックでラップしてください：

```typescript
try {
  const session = await client.createSession(request);
  console.log('成功:', session);
} catch (error) {
  console.error('エラー:', error.message);
}
```

## TypeScript サポート

このライブラリは TypeScript で書かれ、完全な型定義を含んでいます。IntelliSense サポートとコンパイル時型チェックが優れています。

## 貢献

貢献を歓迎します！詳細は[コントリビューションガイドライン](../CONTRIBUTING.md)をお読みください。

## ライセンス

このプロジェクトは MIT ライセンスの下でライセンスされています - 詳細は[LICENSE](../LICENSE) ファイルをご覧ください。

## サポート

サポートについては、[support@jules.ai](mailto:support@jules.ai) までお問い合わせいただくか、[当社のドキュメント](https://developers.google.com/jules/api)をご覧ください。
