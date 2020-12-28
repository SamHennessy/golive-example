package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/SamHennessy/golive-example/internal/components"
	"github.com/SamHennessy/golive-example/internal/domain"
	"github.com/SamHennessy/golive-example/internal/middleware"
	"github.com/brendonferreira/golive"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	ds := domain.NewService()
	liveServer := golive.NewServer()

	var livePage golive.PageContent
	livePage.Lang = "us"
	livePage.Title = "Example Application for GoLive"

	headFileBytes, err := ioutil.ReadFile("templates/head.gohtml")
	if err != nil {
		fmt.Println("ERROR: read head template file:", err)

		os.Exit(1)
	}

	// nolint:gosec // no user input
	livePage.Head = template.HTML(headFileBytes)

	app := fiber.New()
	app.Static("/", "./public")
	app.Get("/", liveServer.CreateHTMLHandlerWithMiddleware(components.NewHome(ds), livePage, middleware.Session))
	app.Get("/ws", websocket.New(liveServer.HandleWSRequest))

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("ERROR: ", err)

		os.Exit(1)
	}
}
