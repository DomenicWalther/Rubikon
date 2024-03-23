package main

import (
	"github.com/domenicwalther/rubikon/backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Post("/Users", handlers.CreateUser)
}
