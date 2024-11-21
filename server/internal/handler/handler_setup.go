package handler

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/middleware"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RouteSetUp(app *fiber.App, serviceImp *services.UserServicesImpl) {
	middleware.SetupMiddleware(app)

	app.Post(
		"/signup",
		serviceImp.SignUp,
	)
	app.Post(
		"/signin",
		serviceImp.SignIn,
	)

	protected := app.Group("/protected", JWTMiddleware("secret"))

	protected.Get(
		"/", func(ctx *fiber.Ctx) error {
			return ctx.SendString("Welcome to Stock Aggregation API")
		})

}

func JWTMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
		}

		tokenString := authHeader[len("Bearer "):]
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
