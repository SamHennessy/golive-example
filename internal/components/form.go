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

	InputText string
	Tasks     []domain.ToDoTask
	ds        *domain.Service
	sessionID string
}

func NewForm(ds *domain.Service) func(ctx context.Context) *golive.LiveComponent {
	return func(ctx context.Context) *golive.LiveComponent {
		var c Form
		c.ds = ds
		c.sessionID = middleware.GetSessionID(ctx)
		c.Tasks = ds.GetToDoList(c.sessionID)
		// f.Tasks = []Task{
		// 	{
		// 		Done: true,
		// 		Text: "Wake up",
		// 	},
		// 	{
		// 		Done: true,
		// 		Text: "Breath",
		// 	},
		// 	{
		// 		Done: false,
		// 		Text: "Turn on the coffee maker",
		// 	},
		// }

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

// BeforeMount the component loading html
func (c *Form) BeforeMount(_ *golive.LiveComponent) {
	// fmt.Println("Form BeforeMount")
}

// Mounted after component is mounted
func (c *Form) Mounted(_ *golive.LiveComponent) {
	// fmt.Println("Form Mounted")
	c.ds.SetToDoList(c.sessionID, c.Tasks)
}

func (c *Form) CanAdd() bool {
	fmt.Println("CanAdd", len(c.InputText) > 0)
	return len(c.InputText) > 0
}

func (c *Form) HandleAdd() {
	if len(c.InputText) > 0 {
		c.Tasks = append(c.Tasks, domain.ToDoTask{
			Done: false,
			Text: c.InputText,
		})
		c.InputText = ""
	}
}
