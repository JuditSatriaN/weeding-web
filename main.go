package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/weeding-web/internal/entity"
	"github.com/weeding-web/routes"
)

func main() {
	engine := html.New(entity.TemplateDirectory, ".html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views:        engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		AppName:      entity.APPName,
	})

	// Setting basic configuration
	app.Use(logger.New(), recover.New())
	app.Static("/static", entity.StaticDirectory)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	routes.BuildCompanyProfileRoutes(app)

	if err := app.Listen(":9000"); err != nil {
		panic(err)
	}
}
