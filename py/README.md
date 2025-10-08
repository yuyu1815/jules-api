# Jules API Python Client

[![PyPI version](https://badge.fury.io/py/jules-api.svg)](https://pypi.org/project/jules-api/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Python Versions](https://img.shields.io/pypi/pyversions/jules-api)
Apology: We apologize for mistakenly presenting this library as official in earlier versions. This is an unofficial library.

Unofficial Python client library for the Jules API.

## Installation

```bash
pip install jules-api
```

## Usage

```python
from jules_api import JulesClient
from jules_api import CreateSessionRequest, SourceContext, GithubRepoContext, SendMessageRequest

# Initialize the client with your API key
client = JulesClient(api_key="YOUR_API_KEY_HERE")

# List available sources
sources_response = client.list_sources()
print("Available sources:")
for source in sources_response.sources:
    print(f"- {source.id}: {source.name}")
    if source.github_repo:
        print(f"  GitHub: {source.github_repo.owner}/{source.github_repo.repo}")

if not sources_response.sources:
    print("No sources found. Please connect a GitHub repository in the Jules web app first.")
    exit(1)

# Create a new session
first_source = sources_response.sources[0]
session_request = CreateSessionRequest(
    prompt="Create a simple web app that displays 'Hello from Jules!'",
    source_context=SourceContext(
        source=first_source.name,
        github_repo_context=GithubRepoContext(starting_branch="main")
    ),
    title="Hello World App Session"
)

session = client.create_session(session_request)
print(f"Created session: {session.id}")
print(f"Title: {session.title}")
print(f"Prompt: {session.prompt}")

# Send a message to the agent
message_request = SendMessageRequest(
    prompt="Please add some styling to make it look more attractive."
)
client.send_message(session.id, message_request)
print("Message sent to the agent!")

# List activities (optional)
activities_response = client.list_activities(session.id, page_size=10)
print(f"\nFound {len(activities_response.activities)} activities:")
for activity in activities_response.activities:
    content = activity.content or "No content"
    if len(content) > 100:
        content = content[:100] + "..."
    print(f"- {activity.type}: {content}")
```

## Alternative Client Creation

You can also use the convenience function:

```python
from jules_api import create_client

client = create_client("YOUR_API_KEY_HERE")
```

## API Reference

### JulesClient

#### Constructor

```python
client = JulesClient(api_key="your_api_key", base_url="https://jules.googleapis.com/v1alpha")
```

**Parameters:**
- `api_key`: Your Jules API key (required)
- `base_url`: API base URL (optional, defaults to "https://jules.googleapis.com/v1alpha")

#### Methods

##### `list_sources(next_page_token=None)`

List all available sources connected to your Jules account.

**Parameters:**
- `next_page_token`: Token for pagination (optional)

**Returns:** `ListSourcesResponse` object with sources array and optional next_page_token

##### `create_session(request)`

Create a new session.

**Parameters:**
- `request`: `CreateSessionRequest` object with session parameters

**Returns:** `Session` object representing the created session

##### `list_sessions(page_size=None, next_page_token=None)`

List your sessions.

**Parameters:**
- `page_size`: Maximum number of sessions to return (optional)
- `next_page_token`: Token for pagination (optional)

**Returns:** `ListSessionsResponse` object with sessions array and optional next_page_token

##### `approve_plan(session_id)`

Approve the latest plan for a session (use when require_plan_approval is True).

**Parameters:**
- `session_id`: The session ID string (required)

##### `list_activities(session_id, page_size=None, next_page_token=None)`

List activities for a session.

**Parameters:**
- `session_id`: The session ID string (required)
- `page_size`: Maximum number of activities to return (optional)
- `next_page_token`: Token for pagination (optional)

**Returns:** `ListActivitiesResponse` object with activities array and optional next_page_token

##### `send_message(session_id, request)`

Send a message to the agent in a session.

**Parameters:**
- `session_id`: The session ID string (required)
- `request`: `SendMessageRequest` object with the message

##### `get_session(session_id)`

Get details of a specific session.

**Parameters:**
- `session_id`: The session ID string (required)

**Returns:** `Session` object

##### `get_source(source_id)`

Get details of a specific source.

**Parameters:**
- `source_id`: The source ID string (required)

**Returns:** `Source` object

## Models

### Request/Response Models

- `CreateSessionRequest`: Used to create new sessions
- `SendMessageRequest`: Used to send messages to agents
- `ListSourcesResponse`: Response from listing sources
- `ListSessionsResponse`: Response from listing sessions
- `ListActivitiesResponse`: Response from listing activities
- `Session`: Session information
- `Source`: Source information
- `Activity`: Activity information

### Data Models

- `SourceContext`: Context for session sources
- `GithubRepoContext`: Additional context for GitHub repositories
- `GithubRepo`: GitHub repository information

All models use Pydantic for validation and type hints.

## Authentication

All API requests require authentication using your Jules API key. Get your API key from the Settings page in the Jules web app. The client automatically includes the key in the `X-Goog-Api-Key` header for all requests.

**Important:** Keep your API keys secure. Never commit them to version control or expose them in client-side code.

Set your API key as an environment variable:

```bash
export JULES_API_KEY=your_api_key_here
```

## Error Handling

The client raises `requests.HTTPError` exceptions for HTTP errors. Always wrap API calls in try-except blocks:

```python
from requests.exceptions import HTTPError

try:
    session = client.create_session(request)
    print("Success:", session)
except HTTPError as e:
    print(f"HTTP Error: {e}")
    print(f"Status Code: {e.response.status_code}")
    print(f"Response: {e.response.text}")
```

## Type Hints

This library uses modern Python type hints throughout. Your IDE should provide excellent autocomplete and type checking support.

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](../CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.
