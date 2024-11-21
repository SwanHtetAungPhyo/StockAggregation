package handler

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/config"
	jwtware "github.com/SwanHtetAungPhyo/stockAggregation/internal/jwt"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/middleware"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/repo"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ServicesMapper struct {
	UserService      *services.UserServicesImpl
	WatchListService *services.WatchListServiceImpl
}

var (
	userRepo         = repo.UserRepo{}
	watchRepo        = repo.NewWatchRepo()
	userService      = services.NewUserServicesImpl(&userRepo)
	watchListService = services.NewWatchListService(watchRepo)

	servicesMapper = &ServicesMapper{
		UserService:      userService,
		WatchListService: watchListService,
	}
)

func RouteSetUp(app *fiber.App) {
	envVar := config.GetEnv()
	middleware.SetupMiddleware(app)

	app.Post(
		"/signup",
		servicesMapper.UserService.SignUp,
	)
	app.Post(
		"/signin",
		servicesMapper.UserService.SignIn,
	)

	protected := app.Group("/protected", jwtware.JWTMiddleware(envVar.JwtSecret))

	protected.Get(
		"/", func(ctx *fiber.Ctx) error {
			return ctx.SendString("Welcome to Stock Aggregation API")
		})

	protected.Post("/watchlist/:id", servicesMapper.WatchListService.AddWatchList)

}
