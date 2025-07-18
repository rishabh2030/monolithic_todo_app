package service

import (
	"todo/internal/models"
	"todo/internal/repository"
)

type TodoService interface {
	CreateTodo(todo *models.Todos) error
	GetTodoByID(id string) (*models.Todos, error)
	GetTodosByUserID(userID string) ([]models.Todos, error)
	GetTodosList(userID string, page, pageSize int) ([]models.Todos, int64, error)
	UpdateTodo(todo *models.Todos) error
	DeleteTodo(id string) error
}

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (s *todoService) CreateTodo(todo *models.Todos) error {
	return s.todoRepo.CreateTodo(todo)
}

func (s *todoService) GetTodoByID(id string) (*models.Todos, error) {
	return s.todoRepo.GetTodoByID(id)
}

func (s *todoService) GetTodosByUserID(userID string) ([]models.Todos, error) {
	return s.todoRepo.GetTodosByUserID(userID)
}

func (s *todoService) GetTodosList(userID string, page, pageSize int) ([]models.Todos, int64, error) {
	return s.todoRepo.GetTodosList(userID, page, pageSize)
}

func (s *todoService) UpdateTodo(todo *models.Todos) error {
	return s.todoRepo.UpdateTodo(todo)
}

func (s *todoService) DeleteTodo(id string) error {
	return s.todoRepo.DeleteTodo(id)
}
