# Jules API JavaScript Client

[![npm version](https://badge.fury.io/js/%40jules-ai%2Fjules-api.svg)](https://badge.fury.io/js/%40jules-ai%2Fjules-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Unofficial Node.js/TypeScript client library for the Jules API.

## Installation

```bash
npm install @yuzumican/jules-api
```

## Usage

```typescript
import { JulesClient } from '@yuzumican/jules-api';

// Initialize the client with your API key
const client = new JulesClient({
  apiKey: 'YOUR_API_KEY_HERE'
});

// List available sources
const sources = await client.listSources();
console.log('Available sources:', sources.sources);

// Create a new session
const session = await client.createSession({
  prompt: 'Create a boba app!',
  sourceContext: {
    source: 'sources/github/yourusername/yourrepo',
    githubRepoContext: {
      startingBranch: 'main'
    }
  },
  title: 'Boba App Session'
});
console.log('Created session:', session);

// Send a message to the agent
await client.sendMessage(session.id, {
  prompt: 'Can you make the app corgi themed?'
});

// List activities in the session
const activities = await client.listActivities(session.id, 30);
console.log('Session activities:', activities.activities);
```

## API Reference

### JulesClient

#### Constructor

```typescript
const client = new JulesClient(options: JulesClientOptions)
```

Options:
- `apiKey`: Your Jules API key (required)
- `baseUrl`: API base URL (optional, defaults to 'https://jules.googleapis.com/v1alpha')

#### Methods

##### `listSources(nextPageToken?: string): Promise<ListSourcesResponse>`

List all available sources connected to your Jules account.

**Parameters:**
- `nextPageToken`: Token for pagination (optional)

**Returns:** Object containing array of sources and optional nextPageToken

##### `createSession(request: CreateSessionRequest): Promise<Session>`

Create a new session.

**Parameters:**
- `request`: Session creation parameters

**CreateSessionRequest:**
- `prompt`: Initial prompt for the session (required)
- `sourceContext`: Source context with<source name and github repo details (required)
- `title`: Session title (required)
- `requirePlanApproval`: Whether to require explicit plan approval (optional, defaults to false)

**Returns:** Created session object

##### `listSessions(pageSize?: number, nextPageToken?: string): Promise<ListSessionsResponse>`

List your sessions.

**Parameters:**
- `pageSize`: Maximum number of sessions to return (optional)
- `nextPageToken`: Token for pagination (optional)

**Returns:** Object containing array of sessions and optional nextPageToken

##### `approvePlan(sessionId: string): Promise<void>`

Approve the latest plan for a session (use when requirePlanApproval is true).

**Parameters:**
- `sessionId`: The session ID (required)

##### `listActivities(sessionId: string, pageSize?: number, nextPageToken?: string): Promise<ListActivitiesResponse>`

List activities for a session.

**Parameters:**
- `sessionId`: The session ID (required)
- `pageSize`: Maximum number of activities to return (optional)
- `nextPageToken`: Token for pagination (optional)

**Returns:** Object containing array of activities and optional nextPageToken

##### `sendMessage(sessionId: string, request: SendMessageRequest): Promise<void>`

Send a message to the agent in a session.

**Parameters:**
- `sessionId`: The session ID (required)
- `request`: Message parameters

**SendMessageRequest:**
- `prompt`: The message to send (required)

##### `getSession(sessionId: string): Promise<Session>`

Get details of a specific session.

**Parameters:**
- `sessionId`: The session ID (required)

**Returns:** Session object

##### `getSource(sourceId: string): Promise<Source>`

Get details of a specific source.

**Parameters:**
- `sourceId`: The source ID (required)

**Returns:** Source object

## Authentication

All API requests require authentication using your Jules API key. Get your API key from the Settings page in the Jules web app. The client automatically includes the key in the `X-Goog-Api-Key` header for all requests.

**Important:** Keep your API keys secure. Never commit them to version control or expose them in client-side code.

## Error Handling

The client throws exceptions for HTTP errors. Always wrap API calls in try-catch blocks:

```typescript
try {
  const session = await client.createSession(request);
  console.log('Success:', session);
} catch (error) {
  console.error('Error:', error.message);
}
```

## TypeScript Support

This library is written in TypeScript and includes full type definitions. You get excellent IntelliSense support and compile-time type checking.

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](../CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.

## Support

For support, please contact [support@jules.ai](mailto:support@jules.ai) or visit our [documentation](https://developers.google.com/jules/api).
