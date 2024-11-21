package services

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/models"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type (
	Service interface {
		SignUp(ctx *fiber.Ctx) error
		SignIn(ctx *fiber.Ctx) error
	}
	UserServicesImpl struct {
		UserRepo *repo.UserRepo
	}
	LoginResponse struct {
		Token string `json:"token"`
		Id    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
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
	id, name, err := u.UserRepo.Login(login)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, "Failed to login user")
	}

	claims := jwt.MapClaims{
		"id":   id,
		"name": login.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create token"})
	}

	return ctx.Status(fiber.StatusOK).JSON(LoginResponse{
		Token: t,
		Id:    id,
		Name:  name,
		Email: login.Email,
	})
}
