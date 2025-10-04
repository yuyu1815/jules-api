package jules

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Client represents a Jules API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	timeout    time.Duration
}

// NewClient creates a new Jules API client
func NewClient(options *ClientOptions) (*Client, error) {
	apiKey := ""
	if options.APIKey != nil {
		apiKey = *options.APIKey
	} else {
		apiKey = os.Getenv("JULES_API_KEY")
	}

	if apiKey == "" {
		return nil, fmt.Errorf("API key must be provided or set as JULES_API_KEY environment variable")
	}

	baseURL := "https://jules.googleapis.com/v1alpha"
	if options.BaseURL != nil {
		baseURL = *options.BaseURL
	}

	timeout := 60 * time.Second
	if options.Timeout != nil {
		timeout = *options.Timeout
	}

	return &Client{
		httpClient: &http.Client{},
		baseURL:    baseURL,
		apiKey:     apiKey,
		timeout:    timeout,
	}, nil
}

// _makeRequest performs an HTTP request and unmarshals the response
func (c *Client) _makeRequest(ctx context.Context, method, endpoint string, body interface{}, result interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("X-Goog-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Indefinite retry logic
	if c.timeout == -1 {
		for {
			// Create a new context for each retry attempt to avoid deadline issues
			reqCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

			reqClone := req.Clone(reqCtx)
			resp, err := c.httpClient.Do(reqClone)

			success := err == nil && resp != nil && resp.StatusCode >= 200 && resp.StatusCode < 300

			if success {
				if result != nil {
					err = json.NewDecoder(resp.Body).Decode(result)
				}
				resp.Body.Close()
				cancel()
				return err
			}

			if resp != nil && resp.StatusCode == http.StatusNotFound {
				fmt.Println("Resource not found (404), retrying in 1 second...")
			} else if err != nil {
				fmt.Printf("Request failed with %v, retrying in 1 second...\n", err)
			} else if resp != nil {
				fmt.Printf("Request failed with status %d, retrying in 1 second...\n", resp.StatusCode)
			}

			if resp != nil {
				resp.Body.Close()
			}
			cancel()

			select {
			case <-time.After(1 * time.Second):
				// continue
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	// Normal request with timeout
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

// ListSources returns a list of available sources
func (c *Client) ListSources(nextPageToken string, timeout *time.Duration) (*ListSourcesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout(timeout))
	defer cancel()

	endpoint := "/sources"
	if nextPageToken != "" {
		params := url.Values{}
		params.Add("nextPageToken", nextPageToken)
		endpoint += "?" + params.Encode()
	}

	var result ListSourcesResponse
	err := c._makeRequest(ctx, "GET", endpoint, nil, &result)
	return &result, err
}

// CreateSession creates a new session
func (c *Client) CreateSession(request *CreateSessionRequest, timeout *time.Duration) (*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout(timeout))
	defer cancel()

	var result Session
	err := c._makeRequest(ctx, "POST", "/sessions", request, &result)
	return &result, err
}

// ListSessions returns a list of sessions
func (c *Client) ListSessions(pageSize int, nextPageToken string, timeout *time.Duration) (*ListSessionsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout(timeout))
	defer cancel()

	params := url.Values{}
	if pageSize > 0 {
		params.Add("pageSize", fmt.Sprintf("%d", pageSize))
	}
	if nextPageToken != "" {
		params.Add("nextPageToken", nextPageToken)
	}

	endpoint := "/sessions"
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	var result ListSessionsResponse
	err := c._makeRequest(ctx, "GET", endpoint, nil, &result)
	return &result, err
}

// ApprovePlan approves the latest plan for a session
func (c *Client) ApprovePlan(sessionID string, timeout *time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout(timeout))
	defer cancel()

	endpoint := fmt.Sprintf("/sessions/%s:approvePlan", sessionID)
	return c._makeRequest(ctx, "POST", endpoint, nil, nil)
}

// ListActivities returns activities for a session
func (c *Client) ListActivities(sessionID string, pageSize int, nextPageToken string, timeout *time.Duration) (*ListActivitiesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout(timeout))
	defer cancel()

	params := url.Values{}
	if pageSize > 0 {
		params.Add("pageSize", fmt.Sprintf("%d", pageSize))
	}
	if nextPageToken != "" {
		params.Add("nextPageToken", nextPageToken)
	}

	endpoint := fmt.Sprintf("/sessions/%s/activities", sessionID)
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	var result ListActivitiesResponse
	err := c._makeRequest(ctx, "GET", endpoint, nil, &result)
	return &result, err
}

// SendMessage sends a message to the agent in a session
func (c *Client) SendMessage(sessionID string, request *SendMessageRequest, timeout *time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout(timeout))
	defer cancel()

	endpoint := fmt.Sprintf("/sessions/%s:sendMessage", sessionID)
	return c._makeRequest(ctx, "POST", endpoint, request, nil)
}

// GetSession retrieves a specific session
func (c *Client) GetSession(sessionID string, timeout *time.Duration) (*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout(timeout))
	defer cancel()

	endpoint := fmt.Sprintf("/sessions/%s", sessionID)
	var result Session
	err := c._makeRequest(ctx, "GET", endpoint, nil, &result)
	return &result, err
}

// GetSource retrieves a specific source
func (c *Client) GetSource(sourceID string, timeout *time.Duration) (*Source, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout(timeout))
	defer cancel()

	endpoint := fmt.Sprintf("/sources/%s", sourceID)
	var result Source
	err := c._makeRequest(ctx, "GET", endpoint, nil, &result)
	return &result, err
}

func (c *Client) getTimeout(t *time.Duration) time.Duration {
	if t != nil {
		return *t
	}
	return c.timeout
}
