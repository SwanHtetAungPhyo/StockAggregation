package services

import "github.com/gofiber/fiber/v2"


func sendError(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"error": message,
	})
}

