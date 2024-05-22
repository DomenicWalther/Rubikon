package main

import (
	"github.com/domenicwalther/rubikon/backend/handlers"
	"github.com/domenicwalther/rubikon/backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Post("/Users", handlers.HandleCreateUser)
	app.Post("/foo", middleware.JWTMiddleware(), func(c *fiber.Ctx) error {
		return c.SendString(c.Locals("sub").(string) + " " + c.Locals("sid").(string))
	})
	app.Post("/Users/process-user-progress", middleware.JWTMiddleware(), handlers.ProcessUserProgress)
	app.Get("/Users/top-users", handlers.HandleGetTopUsers)
	app.Get("/Users/:id", handlers.HandleGetUserById)

	app.Get("/Groups/chat/:id", handlers.GetChatMessages)
	app.Get("/Groups/:id", handlers.HandleGetGroups)
	app.Post("/Groups", middleware.JWTMiddleware(), handlers.HandleCreateGroup)
	app.Post("/Groups/join", middleware.JWTMiddleware(), handlers.JoinGroup)
	app.Delete("/Groups/leave", middleware.JWTMiddleware(), handlers.LeaveGroup)

	app.Post("/Groups/chat/addMessage", middleware.JWTMiddleware(), handlers.CreateNewChatMessage)
}
