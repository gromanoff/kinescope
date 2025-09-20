package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gromanoff/kinescope"
)

func main() {
	ctx := context.Background()
	kc := kinescope.New("YOUR_ACCESS_KEY")

	// Create a demo event
	ev, err := kc.CreateEvent(ctx, kinescope.CreateEventRequest{
		Name:        "Demo",
		Type:        "recurring",
		AutoStart:   true,
		LatencyMode: "standard",
		Scheduled:   &kinescope.Scheduled{Time: time.Now().Add(5 * time.Minute)},
	})
	must(err)
	fmt.Println("created:", ev.ID, ev.StreamKey)

	// Enable event and schedule start
	must(kc.EnableEvent(ctx, ev.ID))
	_, err = kc.ScheduleStream(ctx, ev.ID, time.Now().Add(10*time.Minute))
	must(err)

	// Finish & delete when done (commented)
	// must(kc.FinishEvent(ctx, ev.ID))
	// must(kc.DeleteEvent(ctx, ev.ID))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
