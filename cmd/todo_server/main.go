package main

import (
	"log"
	"todo/internal/config"
	"todo/internal/db"
	"todo/internal/handler"
	"todo/internal/middleware"
	"todo/internal/migrations"
	"todo/internal/repository"
	"todo/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config := config.LoadConfig()

	// Initialize middleware
	middleware.Init(app)
	dbInstance, err := db.InitDB(config.DbEngine, config.Dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Set the DB instance in config
	config.Db = dbInstance

	userRepo := repository.NewUserRepo(dbInstance)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Run migrations after DB is set in config
	migrations.CreateMigrations(config)

	app.Post("/v1/register", userHandler.Register)
	app.Post("/v1/login", userHandler.Login)

	// All other authenticated routes under /v1
	v1 := app.Group("/v1", middleware.JwtMiddleware())

	todoRepo := repository.NewTodoRepo(dbInstance)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	v1.Post("/todos", todoHandler.CreateTodo)
	v1.Get("/todos", todoHandler.GetTodos)
	v1.Get("/todos/:id", todoHandler.GetTodoByID)
	v1.Put("/todos/:id", todoHandler.UpdateTodo)
	v1.Delete("/todos/:id", todoHandler.DeleteTodo)

	app.Listen(":3000")
}
