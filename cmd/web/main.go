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

	logger := golive.NewLoggerBasic()
	logger.Level = golive.LogInfo
	liveServer.Log = logger.Log

	var livePage golive.PageContent
	livePage.Lang = "us"
	livePage.Title = "Example Application for GoLive"

	// baseFileBytes, err := ioutil.ReadFile("./base_html.gohtml")
	// if err != nil {
	// 	panic(fmt.Errorf("ERROR: read head template file: %w", err))
	// }
	//
	// golive.BasePage, err = template.New("BasePage").Parse(string(baseFileBytes))
	// if err != nil {
	// 	panic(fmt.Errorf("ERROR: parse base template file: %w", err))
	// }

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

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		fmt.Println("ERROR: ", err)

		os.Exit(1)
	}
}
