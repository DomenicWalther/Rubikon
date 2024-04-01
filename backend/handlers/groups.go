package handlers

import (
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetGroups(c *fiber.Ctx) error {
	var groups []models.Group
	database.DB.Db.Find(&groups)
	return c.Status(200).JSON(groups)
}
