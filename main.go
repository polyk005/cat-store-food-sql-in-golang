package main

import (
	"cat-food-store/database"
	"cat-food-store/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка подключение к базе данных: %v", err)
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(limiter.New())

	routes.RegisterProductRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
