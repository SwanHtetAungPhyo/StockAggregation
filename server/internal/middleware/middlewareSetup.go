package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,HEAD,PUT,PATCH,POST,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(limiter.New(limiter.Config{
		Max: 5,
		Expiration: 60 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).SendString("Too Many Requests")
		},
	}))
}
