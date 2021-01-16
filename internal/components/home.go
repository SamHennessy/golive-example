package components

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/SamHennessy/golive-example/internal/domain"
	"github.com/brendonmatos/golive"
)

type Home struct {
	golive.LiveComponentWrapper
	Form     *golive.LiveComponent
	ds       *domain.Service
	ShowMenu bool
}

func NewHome(ds *domain.Service) func(ctx context.Context) *golive.LiveComponent {
	newForm := NewForm(ds)

	return func(ctx context.Context) *golive.LiveComponent {
		var c Home
		c.Form = newForm(ctx)

		return golive.NewLiveComponent("Home", &c)
	}
}

func (c *Home) TemplateHandler(_ *golive.LiveComponent) string {
	// This will reread the file on each HTTP request (browser refresh), good for dev but not for production use.
	fileBytes, err := ioutil.ReadFile("templates/home.gohtml")
	if err != nil {
		panic(fmt.Errorf("read template file: %w", err))
	}

	return string(fileBytes)
}

func (c *Home) HandleToggleMenu() {
	c.ShowMenu = !c.ShowMenu
}
