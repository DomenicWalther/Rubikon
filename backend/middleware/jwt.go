package middleware

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type JWTError struct {
	message string
}

func (e JWTError) Error() string {
	return e.message
}
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := extractJWTToken(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token claims are not valid",
			})
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User ID is missing in the token",
			})
		}

		sid, ok := claims["sid"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Session ID is missing in the token",
			})
		}

		c.Locals("sub", sub)
		c.Locals("sid", sid)

		return c.Next()

	}
}

func extractJWTToken(c *fiber.Ctx) (*jwt.Token, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid authorization header")
	}

	jwtToken := parts[1]

	verifiedToken, err := verifyJWTToken(jwtToken)
	if err != nil {
		return nil, err
	}

	return verifiedToken, nil
}
func verifyJWTToken(jwtToken string) (*jwt.Token, error) {
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

	defer jwks.EndBackground()
	token, err := jwt.Parse(jwtToken, jwks.Keyfunc)
	if err != nil {
		cancel()
		return token, errors.New("token is not valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		iat := int64(claims["iat"].(float64))

		currentTime := time.Now().Unix()

		if exp < currentTime {
			cancel()
			return token, errors.New("token has expired")
		}

		if iat > currentTime {
			cancel()
			return token, errors.New("token is not valid yet")
		}
		cancel()
		return token, nil
	} else {
		cancel()
		return token, errors.New("token is not valid")
	}

}
