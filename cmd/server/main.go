package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/yaduvamsi/user-age-api/config"
	"github.com/yaduvamsi/user-age-api/db/sqlc"
	"github.com/yaduvamsi/user-age-api/internal/handler"
	"github.com/yaduvamsi/user-age-api/internal/repository"
	"github.com/yaduvamsi/user-age-api/internal/routes"
	"github.com/yaduvamsi/user-age-api/internal/service"
   "github.com/yaduvamsi/user-age-api/internal/logger"
  	"go.uber.org/zap"
	"github.com/yaduvamsi/user-age-api/internal/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Create Logger
	logg, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer logg.Sync()

	logg.Info("application starting")

	// Connect Database
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logg.Info("database connected")

	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	app.Use(
	middleware.LoggerMiddleware(logg),
)

	routes.SetupRoutes(app, userHandler)

	logg.Info("server started", zap.String("port", "3000"))

	log.Fatal(app.Listen(":3000"))
}