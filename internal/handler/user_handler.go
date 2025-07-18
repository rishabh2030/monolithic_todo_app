package handler

import (
	"todo/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := h.UserService.Register(body.Username, body.Password, body.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	user, err := h.UserService.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login user"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
