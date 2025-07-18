package service

import (
	"errors"
	"log"
	"todo/internal/models"
	"todo/internal/repository"
	"todo/internal/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(username, password, Email string) error
	Login(email, password string) (fiber.Map, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(username, password, Email string) error {
	log.Printf("Starting registration process for email: %s", Email)

	existingUser, err := s.userRepo.GetUserByEmail(Email)
	// If error is not "record not found", then it's a real error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Unexpected error checking for existing user: %v", err)
		return err
	}
	// If user exists (no error and user found), return error
	if existingUser != nil {
		log.Printf("Registration failed: email already exists: %s", Email)
		return errors.New("email already exists")
	}

	log.Printf("Email %s is available, proceeding with registration", Email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    Email,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		log.Printf("Failed to create user in database: %v", err)
		return err
	}

	log.Printf("Successfully registered new user with email: %s", Email)
	return nil
}

func (s *userService) Login(email, password string) (fiber.Map, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return nil, err
	}

	return fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}
