package kinescope

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// CreateEvent -> POST /v2/live/events
func (c *Client) CreateEvent(ctx context.Context, req CreateEventRequest) (*Event, error) {
	var env envelope[Event]
	if err := c.do(ctx, http.MethodPost, "/v2/live/events", req, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}

// UpdateEvent -> PUT /v2/live/events/:event_id
func (c *Client) UpdateEvent(ctx context.Context, eventID string, req UpdateEventRequest) (*Event, error) {
	var env envelope[Event]
	path := fmt.Sprintf("/v2/live/events/%s", eventID)
	if err := c.do(ctx, http.MethodPut, path, req, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}

// RawUpdateEvent sends arbitrary JSON (e.g. {"record": null})
func (c *Client) RawUpdateEvent(ctx context.Context, eventID string, payload map[string]any) (*Event, error) {
	var env envelope[Event]
	path := fmt.Sprintf("/v2/live/events/%s", eventID)
	if err := c.do(ctx, http.MethodPut, path, payload, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}

// GetEvent -> GET /v2/live/events/:event_id
func (c *Client) GetEvent(ctx context.Context, eventID string) (*Event, error) {
	var env envelope[Event]
	path := fmt.Sprintf("/v2/live/events/%s", eventID)
	if err := c.do(ctx, http.MethodGet, path, nil, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}

// EnableEvent -> PUT /v2/live/events/:event_id/enable
func (c *Client) EnableEvent(ctx context.Context, eventID string) error {
	path := fmt.Sprintf("/v2/live/events/%s/enable", eventID)
	return c.do(ctx, http.MethodPut, path, nil, nil)
}

// FinishEvent -> PUT /v2/live/events/:event_id/complete
func (c *Client) FinishEvent(ctx context.Context, eventID string) error {
	path := fmt.Sprintf("/v2/live/events/%s/complete", eventID)
	return c.do(ctx, http.MethodPut, path, nil, nil)
}

// DeleteEvent -> DELETE /v2/live/events/:event_id
func (c *Client) DeleteEvent(ctx context.Context, eventID string) error {
	path := fmt.Sprintf("/v2/live/events/%s", eventID)
	return c.do(ctx, http.MethodDelete, path, nil, nil)
}

// ScheduleStream -> POST /v2/live/events/:event_id/stream
func (c *Client) ScheduleStream(ctx context.Context, eventID string, start time.Time) (*Stream, error) {
	var env envelope[Stream]
	path := fmt.Sprintf("/v2/live/events/%s/stream", eventID)
	body := struct {
		StartedAt string `json:"started_at"`
	}{StartedAt: start.UTC().Format(time.RFC3339Nano)}
	if err := c.do(ctx, http.MethodPost, path, body, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}

// UpdateScheduledStream -> PUT /v2/live/events/:event_id/stream
func (c *Client) UpdateScheduledStream(ctx context.Context, eventID string, start time.Time) (*Stream, error) {
	var env envelope[Stream]
	path := fmt.Sprintf("/v2/live/events/%s/stream", eventID)
	body := struct {
		StartedAt string `json:"started_at"`
	}{StartedAt: start.UTC().Format(time.RFC3339Nano)}
	if err := c.do(ctx, http.MethodPut, path, body, &env); err != nil {
		return nil, err
	}
	return &env.Data, nil
}

// EnableRecording sets folder for recording
func (c *Client) EnableRecording(ctx context.Context, eventID, folderID string) (*Event, error) {
	req := UpdateEventRequest{
		Record: &Record{ParentID: folderID},
	}
	return c.UpdateEvent(ctx, eventID, req)
}

// DisableRecording sets {"record": null}
func (c *Client) DisableRecording(ctx context.Context, eventID string) (*Event, error) {
	return c.RawUpdateEvent(ctx, eventID, map[string]any{"record": nil})
}
