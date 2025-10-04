package jules

import (
	"time"
)

// Source represents an input source for the agent (e.g., a GitHub repository)
type Source struct {
	Name       string      `json:"name"`
	ID         string      `json:"id"`
	GithubRepo *GithubRepo `json:"githubRepo,omitempty"`
}

// GithubRepo contains GitHub repository information
type GithubRepo struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
}

// GithubRepoContext provides additional context for GitHub repositories
type GithubRepoContext struct {
	StartingBranch string `json:"startingBranch,omitempty"`
}

// SourceContext defines the source for a session
type SourceContext struct {
	Source            string              `json:"source"`
	GithubRepoContext *GithubRepoContext  `json:"githubRepoContext,omitempty"`
}

// Session represents a continuous unit of work within a specific context
type Session struct {
	Name          string          `json:"name"`
	ID            string          `json:"id"`
	Title         string          `json:"title"`
	SourceContext *SourceContext  `json:"sourceContext,omitempty"`
	Prompt        string          `json:"prompt,omitempty"`
}

// ListSourcesResponse is the response from listing sources
type ListSourcesResponse struct {
	Sources       []Source `json:"sources"`
	NextPageToken string   `json:"nextPageToken,omitempty"`
}

// CreateSessionRequest is used to create a new session
type CreateSessionRequest struct {
	Prompt             string          `json:"prompt"`
	SourceContext      SourceContext   `json:"sourceContext"`
	Title              string          `json:"title"`
	RequirePlanApproval bool           `json:"requirePlanApproval,omitempty"`
}

// ListSessionsResponse is the response from listing sessions
type ListSessionsResponse struct {
	Sessions      []Session `json:"sessions"`
	NextPageToken string    `json:"nextPageToken,omitempty"`
}

// Activity represents a single unit of work within a Session
type Activity struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Content   string    `json:"content,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// ListActivitiesResponse is the response from listing activities
type ListActivitiesResponse struct {
	Activities    []Activity `json:"activities"`
	NextPageToken string     `json:"nextPageToken,omitempty"`
}

// SendMessageRequest is used to send a message to the agent
type SendMessageRequest struct {
	Prompt string `json:"prompt"`
}

// ClientOptions contains configuration options for the Jules client
type ClientOptions struct {
	APIKey  string
	BaseURL string
}

// NewClientOptions creates default client options
func NewClientOptions(apiKey string) *ClientOptions {
	return &ClientOptions{
		APIKey:  apiKey,
		BaseURL: "https://jules.googleapis.com/v1alpha",
	}
}
