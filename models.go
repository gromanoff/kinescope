package kinescope

import "time"

// envelope for successful responses
type envelope[T any] struct {
	Data T `json:"data"`
}

// ---------- Event ----------

type Event struct {
	ID           string     `json:"id"`
	WorkspaceID  string     `json:"workspace_id"`
	ParentID     string     `json:"parent_id"`
	Name         string     `json:"name"`
	Subtitle     string     `json:"subtitle"`
	Type         string     `json:"type"` // "one-time" | "recurring"
	StreamKey    string     `json:"streamkey"`
	AutoStart    bool       `json:"auto_start"`
	Protected    bool       `json:"protected"`
	TimeShift    bool       `json:"time_shift"`
	ReconnectWin int        `json:"reconnect_window"`
	PlayLink     string     `json:"play_link"`
	RTMPLink     string     `json:"rtmp_link"`
	Scheduled    *Scheduled `json:"scheduled,omitempty"`
	Record       *Record    `json:"record,omitempty"`
	Stream       *Stream    `json:"stream,omitempty"`
	LatencyMode  string     `json:"latency_mode"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type Scheduled struct {
	Time time.Time `json:"time"`
}

type Record struct {
	ParentID string `json:"parent_id"`
}

// ---------- Stream (schedule) ----------

type Stream struct {
	ID        string     `json:"id"`
	EventID   string     `json:"event_id"`
	Status    string     `json:"status"` // "pending" | "running" | ...
	StartedAt time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
}

// ---------- Requests ----------

type CreateEventRequest struct {
	Name            string     `json:"name"`
	Subtitle        string     `json:"subtitle,omitempty"`
	Type            string     `json:"type"`
	AutoStart       bool       `json:"auto_start,omitempty"`
	Protected       bool       `json:"protected,omitempty"`
	TimeShift       bool       `json:"time_shift,omitempty"`
	ParentID        string     `json:"parent_id,omitempty"`
	ReconnectWindow int        `json:"reconnect_window,omitempty"`
	Scheduled       *Scheduled `json:"scheduled,omitempty"`
	Record          *Record    `json:"record,omitempty"`
	LatencyMode     string     `json:"latency_mode,omitempty"`
	Restreams       []Restream `json:"restreams,omitempty"`
}

type UpdateEventRequest struct {
	Name            *string    `json:"name,omitempty"`
	Subtitle        *string    `json:"subtitle,omitempty"`
	AutoStart       *bool      `json:"auto_start,omitempty"`
	Protected       *bool      `json:"protected,omitempty"`
	TimeShift       *bool      `json:"time_shift,omitempty"`
	ReconnectWindow *int       `json:"reconnect_window,omitempty"`
	Scheduled       *Scheduled `json:"scheduled,omitempty"`
	Record          *Record    `json:"record,omitempty"` // to set null use RawUpdateEvent
	LatencyMode     *string    `json:"latency_mode,omitempty"`
	Moderators      []string   `json:"moderators,omitempty"`
	ShowMembers     *bool      `json:"show_members,omitempty"`
	ChatPreview     *bool      `json:"chat_preview,omitempty"`
}

type Restream struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
	Enabled     *bool  `json:"enabled,omitempty"`
}
