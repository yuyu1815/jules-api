package jules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client represents a Jules API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// NewClient creates a new Jules API client
func NewClient(options *ClientOptions) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    options.BaseURL,
		apiKey:     options.APIKey,
	}
}

// newRequest creates a new HTTP request with proper headers
func (c *Client) newRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	fullURL := c.baseURL + endpoint
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Goog-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// doRequest performs an HTTP request and unmarshals the response
func (c *Client) doRequest(method, endpoint string, body interface{}, result interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := c.newRequest(method, endpoint, bodyReader)
	if err != nil {
		return err
	}

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
func (c *Client) ListSources(nextPageToken string) (*ListSourcesResponse, error) {
	endpoint := "/sources"
	if nextPageToken != "" {
		params := url.Values{}
		params.Add("nextPageToken", nextPageToken)
		endpoint += "?" + params.Encode()
	}

	var result ListSourcesResponse
	err := c.doRequest("GET", endpoint, nil, &result)
	return &result, err
}

// CreateSession creates a new session
func (c *Client) CreateSession(request *CreateSessionRequest) (*Session, error) {
	var result Session
	err := c.doRequest("POST", "/sessions", request, &result)
	return &result, err
}

// ListSessions returns a list of sessions
func (c *Client) ListSessions(pageSize int, nextPageToken string) (*ListSessionsResponse, error) {
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
	err := c.doRequest("GET", endpoint, nil, &result)
	return &result, err
}

// ApprovePlan approves the latest plan for a session
func (c *Client) ApprovePlan(sessionID string) error {
	endpoint := fmt.Sprintf("/sessions/%s:approvePlan", sessionID)
	return c.doRequest("POST", endpoint, nil, nil)
}

// ListActivities returns activities for a session
func (c *Client) ListActivities(sessionID string, pageSize int, nextPageToken string) (*ListActivitiesResponse, error) {
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
	err := c.doRequest("GET", endpoint, nil, &result)
	return &result, err
}

// SendMessage sends a message to the agent in a session
func (c *Client) SendMessage(sessionID string, request *SendMessageRequest) error {
	endpoint := fmt.Sprintf("/sessions/%s:sendMessage", sessionID)
	return c.doRequest("POST", endpoint, request, nil)
}

// GetSession retrieves a specific session
func (c *Client) GetSession(sessionID string) (*Session, error) {
	endpoint := fmt.Sprintf("/sessions/%s", sessionID)
	var result Session
	err := c.doRequest("GET", endpoint, nil, &result)
	return &result, err
}

// GetSource retrieves a specific source
func (c *Client) GetSource(sourceID string) (*Source, error) {
	endpoint := fmt.Sprintf("/sources/%s", sourceID)
	var result Source
	err := c.doRequest("GET", endpoint, nil, &result)
	return &result, err
}
