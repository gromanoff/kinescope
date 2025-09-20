# Kinescope Go SDK (Unofficial, typed, minimal)

Typed, context-aware SDK for Kinescope Live Events:
- Create/Update/Delete event
- Enable event (start allowed)
- Finish event
- Schedule / Update schedule
- Get event
- Recording helpers (enable/disable)

No `map[string]any` – only typed requests and responses. Proper error type with Kinescope error code/message. Small, dependency-free.

## Install

```bash
go get github.com/gromanoff/kinescope
```

## Quick Start

```go
package main

import (
  "context"
  "fmt"
  "time"

  "github.com/gromanoff/kinescope"
)

func main() {
  ctx := context.Background()
  kc := kinescope.New("YOUR_ACCESS_KEY") // defaults to https://api.kinescope.io

  // 1) Create event
  ev, err := kc.CreateEvent(ctx, kinescope.CreateEventRequest{
    Name:        "New event",
    Type:        "recurring",
    AutoStart:   true,
    LatencyMode: "standard",
    Record:      &kinescope.Record{ParentID: "c21d86ac-7e90-43e7-b825-cbf300951355"},
    Scheduled:   &kinescope.Scheduled{Time: time.Now().Add(10 * time.Minute)},
  })
  must(err)
  fmt.Println("event id:", ev.ID, "streamkey:", ev.StreamKey)

  // 2) Enable event
  must(kc.EnableEvent(ctx, ev.ID))

  // 3) Schedule start
  _, err = kc.ScheduleStream(ctx, ev.ID, time.Now().Add(15*time.Minute))
  must(err)

  // 4) Disable recording (example)
  _, _ = kc.DisableRecording(ctx, ev.ID)

  // 5) Finish (when done)
  // must(kc.FinishEvent(ctx, ev.ID))
}

func must(err error) {
  if err != nil {
    panic(err)
  }
}
```

## Auth / Student access

* `protected: false` → anyone can watch via `play_link`/embed – **no extra auth** needed.
* `protected: true` → set up an authorization backend and issue signed tokens to the player.

## License

MIT
