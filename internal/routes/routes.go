package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yaduvamsi/user-age-api/internal/handler"
)

func SetupRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
) {
	app.Get("/users", userHandler.ListUsers)
	app.Get("/users/:id", userHandler.GetUser)
	app.Post("/users", userHandler.CreateUser)

}