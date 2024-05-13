package handlers

import (
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetChatMessages(c *fiber.Ctx) error {
	groupID := c.Params("id")
	var ChatMessages []models.GroupChatMessage
	database.DB.Db.Where("group_id = ?", groupID).Order("created_at DESC").Limit(20).Find(&ChatMessages)

	return c.Status(200).JSON(ChatMessages)
}

func CreateNewChatMessage(c *fiber.Ctx) error {
	message := new(models.GroupChatMessage)
	if err := c.BodyParser(message); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}
	message.UserID = c.Locals("sub").(string)
	database.DB.Db.Create(&message)

	return c.Status(200).JSON(message)
}
