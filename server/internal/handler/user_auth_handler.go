package  handler

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RouteSetUp(app *fiber.App)  {
	app.Post(
		"/signup",
		services.SignUp,
	)
	app.Post(
		"/signin",
		services.SignIn,
	)
}