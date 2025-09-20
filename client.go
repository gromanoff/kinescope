package kinescope

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const DefaultBaseURL = "https://api.kinescope.io"

// Client is the Kinescope API client. Configure with a bearer token and optional settings.
type Client struct {
	BaseURL   string
	HTTP      *http.Client
	AuthToken string
	UserAgent string
}

// Option allows configuring Client with functional options.
type Option func(*Client)

// New returns a new Client with the given bearer token and options.
func New(token string, opts ...Option) *Client {
	c := &Client{
		BaseURL:   DefaultBaseURL,
		HTTP:      &http.Client{Timeout: 15 * time.Second},
		AuthToken: token,
		UserAgent: "kinescope-go/0.2",
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithBaseURL sets the base API URL.
func WithBaseURL(u string) Option { return func(c *Client) { c.BaseURL = u } }

// WithHTTP sets a custom HTTP client.
func WithHTTP(h *http.Client) Option { return func(c *Client) { c.HTTP = h } }

// WithUserAgent sets a custom User-Agent header.
func WithUserAgent(ua string) Option { return func(c *Client) { c.UserAgent = ua } }

// do performs an HTTP request with JSON marshaling/unmarshaling and auth headers.
func (c *Client) do(ctx context.Context, method, path string, in any, out any) error {
	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}

	u, err := url.JoinPath(c.BaseURL, path)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, u, body)
	if err != nil {
		return err
	}
	if c.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		return parseAPIError(b)
	}

	if out == nil || len(b) == 0 {
		return nil
	}
	return json.Unmarshal(b, out)
}
