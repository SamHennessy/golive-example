package components

import (
	"context"
	"sync"
	"time"

	"github.com/brendonferreira/golive"
)

type Clock struct {
	golive.LiveComponentWrapper
	ActualTime time.Time
	mountOnce  sync.Once
	ticker     *time.Ticker
	done       chan bool
	Pause      bool
}

func NewClock(_ context.Context) *golive.LiveComponent {
	return golive.NewLiveComponent("Clock", &Clock{
		ActualTime: time.Now(),
	})
}

func (c *Clock) Mounted(_ *golive.LiveComponent) {
	c.mountOnce.Do(func() {
		c.ticker = time.NewTicker(time.Second)
		c.done = make(chan bool)

		go func() {
			for {
				select {
				case <-c.done:
					return
				case t := <-c.ticker.C:
					if !c.Pause {
						c.ActualTime = t
						go c.Commit()
					}
				}
			}
		}()
	})
}

func (c *Clock) BeforeUnmount(lc *golive.LiveComponent) {
	c.done <- true
}

func (c *Clock) TemplateHandler(_ *golive.LiveComponent) string {
	return `
<div class="columns">
    <div class="column is-1">
		<span class="is-size-3 is-uppercase has-text-weight-bold">{{ .DayOfMonth }}</span>
	</div>
    <div class="column">
		<span class="has-text-weight-bold">{{ .Month }}</span><br>
		<span class="has-text-weight-light">{{ .Year }}</span>
    </div>
    <div class="column has-text-right">
		<span class="is-size-5 is-uppercase has-text-weight-semibold">{{ .DayOfWeek }}</span><br>
		<span class="is-family-code has-text-weight-light">{{ .Time }}</span>
		<a go-live-click="PlayPause"><i class="mdi mdi-{{ if .Pause }}play{{ else }}pause{{ end }}"></i></a>
	</div>
</div>
	`
}

func (c *Clock) DayOfMonth() string {
	return c.ActualTime.Format("02")
}

func (c *Clock) DayOfWeek() string {
	return c.ActualTime.Format("Monday")
}

func (c *Clock) Month() string {
	return c.ActualTime.Format("Jan")
}

func (c *Clock) Year() string {
	return c.ActualTime.Format("2006")
}

func (c *Clock) Time() string {
	return c.ActualTime.Format("15:04:05 MST")
}

func (c *Clock) PlayPause() {
	c.Pause = !c.Pause
}
