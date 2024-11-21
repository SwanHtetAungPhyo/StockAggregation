package handler

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RouteSetUp(app *fiber.App, serviceImp *services.UserServicesImpl) {
	app.Post(
		"/signup",
		serviceImp.SignUp,
	)
	app.Post(
		"/signin",
		serviceImp.SignIn,
	)
}
