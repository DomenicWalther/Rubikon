package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/MicahParks/keyfunc"
	"github.com/domenicwalther/rubikon/backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Post("/Users", handlers.CreateUser)
	app.Post("/foo", randomMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func randomMiddleware(c *fiber.Ctx) error {
	// Do something
	extractJWTToken(c)
	return c.Next()
}

func extractJWTToken(c *fiber.Ctx) (string, error) {
	authHeader := c.GetReqHeaders()["Authorization"]
	parts := strings.Split(authHeader[0], " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "error", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header",
		})
	}

	jwtToken := parts[1]

	verifyJWTToken(jwtToken)

	return jwtToken, nil
}

func verifyJWTToken(jwtToken string) {
	jwksURL := os.Getenv("JWT_URL")

	ctx, cancel := context.WithCancel(context.Background())

	options := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error refreshing the JWKS key: %s", err)
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(jwksURL, options)

	if err != nil {
		log.Fatalf("Failed to create JWKS from the URL: %s", err)
	}

	token, err := jwt.Parse(jwtToken, jwks.Keyfunc)
	if err != nil {
		log.Fatalf("Failed to parse the JWT: %s", err)
	}

	if !token.Valid {
		log.Fatalf("The JWT is invalid")
	}
	log.Println("The JWT is valid")

	cancel()

	jwks.EndBackground()
}
