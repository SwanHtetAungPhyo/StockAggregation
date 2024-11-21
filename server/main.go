package main

import (
	"encoding/json"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/db"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/handler"
	"github.com/gofiber/fiber/v2"

	"github.com/SwanHtetAungPhyo/stockAggregation/internal/models"
)

func main() {
	db.DbInit()
	db.Migration(models.User{}, models.StockWatchList{})
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	handler.RouteSetUp(app)
	err := app.Listen(":3002")
	if err != nil {
		return
	}
}
