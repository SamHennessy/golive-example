package components

import (
	"context"
	"time"

	"github.com/brendonferreira/golive"
)

type Clock struct {
	golive.LiveComponentWrapper
	ActualTime string
}

func NewClock(ctx context.Context) *golive.LiveComponent {
	return golive.NewLiveComponent("Clock", &Clock{
		ActualTime: "91230192301390193",
	})
}

func (t *Clock) Mounted(_ *golive.LiveComponent) {
	go func() {
		for {
			t.ActualTime = time.Now().Format(time.RFC3339Nano)
			time.Sleep((time.Second * 1) / 60)
			t.Commit()
		}
	}()
}

func (t *Clock) TemplateHandler(_ *golive.LiveComponent) string {
	return `
		<div>
			<span>Time: {{ .ActualTime }}</span>
		</div>
	`
}
