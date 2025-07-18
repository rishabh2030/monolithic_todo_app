package repository

import (
	"log"
	"todo/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}
	log.Printf("Successfully created user with email: %s", user.Email)
	return nil
}

func (r *UserRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("User not found with ID: %s", id)
		} else {
			log.Printf("Error fetching user by ID: %v", err)
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("No existing user found with email: %s (this is normal for new registrations)", email)
		} else {
			log.Printf("Error fetching user by email: %v", err)
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(user *models.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return err
	}
	log.Printf("Successfully updated user with ID: %s", user.ID)
	return nil
}

func (r *UserRepo) DeleteUser(id string) error {
	err := r.db.Delete(&models.User{}, id).Error
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}
	log.Printf("Successfully deleted user with ID: %s", id)
	return nil
}
