package main

import (
	"os"

	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://192.168.0.111:5173/",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_APP_KEY"),
		Secret:  os.Getenv("PUSHER_APP_SECRET"),
		Cluster: "eu",
		Secure:  true,
	}

	setupRoutes(app, pusherClient)

	app.Listen(":3000")
}
