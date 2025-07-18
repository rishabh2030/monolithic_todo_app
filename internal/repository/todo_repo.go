package repository

import (
	"todo/internal/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo *models.Todos) error
	GetTodoByID(id string) (*models.Todos, error)
	GetTodosByUserID(userID string) ([]models.Todos, error)
	GetTodosList(userID string, page, pageSize int) ([]models.Todos, int64, error)
	UpdateTodo(todo *models.Todos) error
	DeleteTodo(id string) error
}

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) TodoRepository {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) CreateTodo(todo *models.Todos) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepo) GetTodoByID(id string) (*models.Todos, error) {
	var todo models.Todos
	if err := r.db.First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepo) GetTodosByUserID(userID string) ([]models.Todos, error) {
	var todos []models.Todos
	if err := r.db.Where("created_by_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepo) GetTodosList(userID string, page, pageSize int) ([]models.Todos, int64, error) {
	var todos []models.Todos
	var total int64

	// Get total count
	if err := r.db.Model(&models.Todos{}).Where("created_by_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	if err := r.db.Where("created_by_id = ?", userID).
		Offset(offset).
		Limit(pageSize).
		Find(&todos).Error; err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}

func (r *TodoRepo) UpdateTodo(todo *models.Todos) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepo) DeleteTodo(id string) error {
	return r.db.Delete(&models.Todos{}, "id = ?", id).Error
}
