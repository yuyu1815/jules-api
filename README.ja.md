# Jules API クライアントライブラリ

これは非公式ライブラリです。

Jules API のクライアントライブラリです。JavaScript/Node.js、Go、Python の 3 言語で提供されています。

## Jules API について

Jules API を使用して、Jules の機能をプログラムから活用し、ソフトウェア開発ライフサイクルを自動化・強化できます。API を使用して、カスタムワークフローを作成したり、コードレビューのようなタスクを自動化したり、日常的に使用するツールに Jules のインテリジェンスを直接埋め込んだりできます。

**注意:** Jules API はアルファリリース段階です。実験的なものです。

## ライブラリ

- [**JavaScript/Node.js**](https://github.com/yuyu1815/jules-api/tree/main/js) - npm パッケージ、TypeScript 対応
- [**Go**](https://github.com/yuyu1815/jules-api/tree/main/go) - Go モジュール
- [**Python**](https://github.com/yuyu1815/jules-api/tree/main/py) - Python パッケージ

各ライブラリは、Jules API と連携するための完全なクライアント実装を提供します：

- API キーを使用した認証
- ソース管理
- セッション管理
- アクティビティ処理
- メッセージ送信

## 使い始め

優先する言語を選択し、それぞれのライブラリの README で指示に従ってください。

## 認証

すべてのリクエストには、Jules Web アプリの Settings ページから取得した API キーが必要です。キーは `X-Goog-Api-Key` ヘッダーで渡してください。

API キーを安全に保管し、パブリックコードに公開しないでください。

## API の概念

- **ソース**: エージェントの入力ソース (GitHub リポジトリなど)
- **セッション**: 特定のコンテキスト内の継続的な作業単位
- **アクティビティ**: セッション内の単一の作業単位

## 重要事項

この API は実験的であり、変更される可能性があります。プロダクションでの使用には安定版のリリースをお待ちください。
