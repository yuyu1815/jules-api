# Jules API Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/yuyu1815/jules-api/go.svg)](https://pkg.go.dev/github.com/yuyu1815/jules-api/go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Apology: We apologize for mistakenly presenting this library as official in earlier versions. This is an unofficial library.

Unofficial Go client library for the Jules API.

## Installation

Run the following command to install the library:

```bash
go get github.com/yuyu1815/jules-api/go@latest
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	jules "github.com/yuyu1815/jules-api/go"
)

func main() {
	// Initialize the client with your API key
	client := jules.NewClient(jules.NewClientOptions(os.Getenv("JULES_API_KEY")))

	// List available sources
	sources, err := client.ListSources("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Available sources:")
	for _, source := range sources.Sources {
		fmt.Printf("- %s: %s\n", source.ID, source.Name)
		if source.GithubRepo != nil {
			fmt.Printf("  GitHub: %s/%s\n", source.GithubRepo.Owner, source.GithubRepo.Repo)
		}
	}

	if len(sources.Sources) == 0 {
		fmt.Println("No sources found. Please connect a GitHub repository in the Jules web app first.")
		return
	}

	// Create a new session
	firstSource := sources.Sources[0]
	session, err := client.CreateSession(&jules.CreateSessionRequest{
		Prompt: "Create a simple web app that displays 'Hello from Jules!'",
		SourceContext: jules.SourceContext{
			Source: firstSource.Name,
			GithubRepoContext: &jules.GithubRepoContext{
				StartingBranch: "main",
			},
		},
		Title: "Hello World App Session",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created session: %s\n", session.ID)
	fmt.Printf("Title: %s\n", session.Title)
	fmt.Printf("Prompt: %s\n", session.Prompt)

	// Send a message to the agent
	err = client.SendMessage(session.ID, &jules.SendMessageRequest{
		Prompt: "Please add some styling to make it look more attractive.",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message sent to the agent!")
}
```

## API Reference

### Client

#### Constructor

```go
func NewClient(options *ClientOptions) *Client
```

Creates a new Jules API client.

**ClientOptions:**
- `APIKey`: Your Jules API key (required)
- `BaseURL`: API base URL (optional, defaults to "https://jules.googleapis.com/v1alpha")

#### Methods

##### `ListSources(nextPageToken string) (*ListSourcesResponse, error)`

List all available sources connected to your Jules account.

**Parameters:**
- `nextPageToken`: Token for pagination (can be empty string)

**Returns:** ListSourcesResponse containing array of sources and optional nextPageToken

##### `CreateSession(request *CreateSessionRequest) (*Session, error)`

Create a new session.

**Parameters:**
- `request`: Session creation parameters

**CreateSessionRequest:**
- `Prompt`: Initial prompt for the session (required)
- `SourceContext`: Source context with source name and GitHub repo details (required)
- `Title`: Session title (required)
- `RequirePlanApproval`: Whether to require explicit plan approval (optional, defaults to false)

**Returns:** Created session object

##### `ListSessions(pageSize int, nextPageToken string) (*ListSessionsResponse, error)`

List your sessions.

**Parameters:**
- `pageSize`: Maximum number of sessions to return (0 for default)
- `nextPageToken`: Token for pagination (can be empty string)

**Returns:** ListSessionsResponse containing array of sessions and optional nextPageToken

##### `ApprovePlan(sessionID string) error`

Approve the latest plan for a session (use when RequirePlanApproval is true).

**Parameters:**
- `sessionID`: The session ID (required)

##### `ListActivities(sessionID string, pageSize int, nextPageToken string) (*ListActivitiesResponse, error)`

List activities for a session.

**Parameters:**
- `sessionID`: The session ID (required)
- `pageSize`: Maximum number of activities to return (0 for default)
- `nextPageToken`: Token for pagination (can be empty string)

**Returns:** ListActivitiesResponse containing array of activities and optional nextPageToken

##### `SendMessage(sessionID string, request *SendMessageRequest) error`

Send a message to the agent in a session.

**Parameters:**
- `sessionID`: The session ID (required)
- `request`: Message parameters

**SendMessageRequest:**
- `Prompt`: The message to send (required)

##### `GetSession(sessionID string) (*Session, error)`

Get details of a specific session.

**Parameters:**
- `sessionID`: The session ID (required)

**Returns:** Session object

##### `GetSource(sourceID string) (*Source, error)`

Get details of a specific source.

**Parameters:**
- `sourceID`: The source ID (required)

**Returns:** Source object

## Authentication

All API requests require authentication using your Jules API key. Get your API key from the Settings page in the Jules web app. The client automatically includes the key in the `X-Goog-Api-Key` header for all requests.

**Important:** Keep your API keys secure. Never commit them to version control or expose them in public code.

Set your API key as an environment variable:

```bash
export JULES_API_KEY=your_api_key_here
```

## Error Handling

The client returns Go errors for failed API requests. Always check for errors:

```go
session, err := client.CreateSession(request)
if err != nil {
	log.Fatal(err)
}
```

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](../CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.
