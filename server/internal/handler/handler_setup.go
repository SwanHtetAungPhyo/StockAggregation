package handler

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/middleware"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RouteSetUp(app *fiber.App, serviceImp *services.UserServicesImpl) {
	middleware.Setup(app)

	app.Post(
		"/signup",
		serviceImp.SignUp,
	)
	app.Post(
		"/signin",
		serviceImp.SignIn,
	)

	app.Get(
		"/", func(ctx *fiber.Ctx) error {
			return ctx.SendString("Welcome to Stock Aggregation API")
		})

}
