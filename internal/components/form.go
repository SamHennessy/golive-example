package components

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/SamHennessy/golive-example/internal/domain"
	"github.com/SamHennessy/golive-example/internal/middleware"
	"github.com/brendonferreira/golive"
)

type Form struct {
	golive.LiveComponentWrapper

	Clock     *golive.LiveComponent
	Tasks     []domain.ToDoTask
	InputText string
	ShowModal bool
	Tab       string
	TaskCount int

	ds        *domain.Service
	sessionID string
}

func NewForm(ds *domain.Service) func(ctx context.Context) *golive.LiveComponent {
	return func(ctx context.Context) *golive.LiveComponent {
		var c Form
		c.ds = ds
		c.sessionID = middleware.GetSessionID(ctx)
		c.Tasks = ds.GetToDoList(c.sessionID)
		c.Tab = "all"
		c.Clock = NewClock(ctx)

		return golive.NewLiveComponent("Form", &c)
	}
}

func (c *Form) TemplateHandler(_ *golive.LiveComponent) string {
	// This will reread the file on each HTTP request (browser refresh), good for dev but not for production use.
	fileBytes, err := ioutil.ReadFile("templates/form.gohtml")
	if err != nil {
		panic(fmt.Errorf("read template file: %w", err))
	}

	return string(fileBytes)
}

// Mounted after component is mounted
func (c *Form) Mounted(_ *golive.LiveComponent) {
	c.ds.SetToDoList(c.sessionID, c.Tasks)
}

func (c *Form) InputInvalid() string {
	if c.InputText != "" && len(c.InputText) < 3 {
		return "Please make it longer"
	}

	return ""
}

func (c *Form) CanAdd() bool {
	return len(c.InputText) > 2
}

func (c *Form) HandleAdd() {
	if len(c.InputText) > 2 {
		c.Tasks = append(c.Tasks, domain.ToDoTask{
			Done: false,
			Text: c.InputText,
		})
		c.InputText = ""
		c.ShowModal = false
	}
}

func (c *Form) HandleShowModal() {
	c.ShowModal = true
}

func (c *Form) HandleHideModal() {
	c.ShowModal = false
}

func (c *Form) ShowTab(data map[string]string) {
	tab := "all"

	switch data["tab"] {
	case "todo":
		tab = "todo"
	case "done":
		tab = "done"
	}

	c.Tab = tab
}

func (c *Form) TabCount() int {
	tc := 0

	for i := 0; i < len(c.Tasks); i++ {
		switch c.Tab {
		case "all":
			tc++
		case "todo":
			if !c.Tasks[i].Done {
				tc++
			}
		case "done":
			if c.Tasks[i].Done {
				tc++
			}
		}
	}

	return tc
}

func (c *Form) ShowTask(task domain.ToDoTask) bool {
	if c.Tab == "todo" && task.Done {
		return false
	}

	if c.Tab == "done" && !task.Done {
		return false
	}

	return true
}
