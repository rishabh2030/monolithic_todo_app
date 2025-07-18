package handler

import (
	"time"
	"todo/internal/models"
	"todo/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todos
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	todo.CreatedById = uuid.MustParse(c.Locals("user_id").(string))
	todo.UpdatedById = uuid.MustParse(c.Locals("user_id").(string))
	h.todoService.CreateTodo(&todo)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Todo created successfully"})
}

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	// Get pagination parameters from query
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)

	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	todos, total, err := h.todoService.GetTodosList(userID, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get todos"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todos,
		"meta": fiber.Map{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

func (h *TodoHandler) GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := h.todoService.GetTodoByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get todo"})
	}
	return c.Status(fiber.StatusOK).JSON(todo)
}

func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("user_id").(string)
	var todo models.Todos
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	oldTodo, err := h.todoService.GetTodoByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get todo"})
	}
	oldTodo.Title = todo.Title
	oldTodo.Description = todo.Description
	oldTodo.UpdatedById = uuid.MustParse(userID)
	oldTodo.UpdatedAt = time.Now()
	err = h.todoService.UpdateTodo(oldTodo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update todo"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Todo updated successfully"})
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	h.todoService.DeleteTodo(id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Todo deleted successfully"})
}
