package main

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/db"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.DbInit()
	db.Migration(models.User{}, models.StockWatchList{})
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(200)
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3002")
	if err != nil {
		return
	}
}
