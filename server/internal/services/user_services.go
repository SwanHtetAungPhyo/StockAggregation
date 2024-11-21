package services

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/models"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/repo"
	"github.com/gofiber/fiber/v2"
)

type (
	Service interface {
		SignUp(ctx *fiber.Ctx) error
		SignIn(ctx *fiber.Ctx) error
	}
	UserServicesImpl struct {
		UserRepo *repo.UserRepo
	}
)

func NewUserServicesImpl(userRepo *repo.UserRepo) *UserServicesImpl {
	return &UserServicesImpl{UserRepo: userRepo}
}

func (u *UserServicesImpl) SignUp(ctx *fiber.Ctx) error {
	var userToBeSaved models.User
	if err := ctx.BodyParser(&userToBeSaved); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, "Invalid input data")
	}

	if err := userToBeSaved.Validate(); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err.Error())
	}

	if err := u.UserRepo.Create(&userToBeSaved); err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, "Failed to create user")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (u *UserServicesImpl) SignIn(ctx *fiber.Ctx) error {
	login := new(models.User)
	if err := ctx.BodyParser(login); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, "Invalid input data")
	}
	if err := login.Validate(); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err.Error())
	}
	if err := u.UserRepo.Login(login); err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, "Failed to login user")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
	})
}
