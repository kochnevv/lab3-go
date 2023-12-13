package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
	"web-currency-parser/internal/controller/handler"
)

func main() {
	engine := html.New("./internal/template", ".html")
	engine.Debug(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	currencyHandler := handler.NewCurrencyHandler()

	v1 := app.Group("/")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	v1.Post("/", currencyHandler.CreateRequest)

	log.Info("Starting http server: localhost:8000")
	if err := app.Listen(fmt.Sprintf(":%d", 8000)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}
}
