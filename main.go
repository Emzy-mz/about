package main

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

type indexPage struct {
	Motd string
}

func main() {
	app := fiber.New()

	app.Use(favicon.New(favicon.Config{
		File: "./public/img/favicon.ico",
	}))

	app.Static("/img", "./public/img")
	app.Static("/css", "./public/css")

	app.Get("/", indexHandler)

	app.Listen(":8080")

}

func indexHandler(c *fiber.Ctx) error {
	p := indexPage{Motd: "Femboys are cute uwu"}
	t, _ := template.ParseFiles("./assets/html/index.html")
	w := new(bytes.Buffer)
	t.Execute(w, p)
	c.Set("Content-type", "text/html; charset=utf-8")
	return c.Send([]byte(w.String()))
}
