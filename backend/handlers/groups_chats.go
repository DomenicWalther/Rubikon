package handlers

import (
	"github.com/domenicwalther/rubikon/backend/data"
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pusher/pusher-http-go/v5"
)

func GetChatMessages(c *fiber.Ctx) error {
	groupID := c.Params("id")
	var ChatMessages []models.GroupChatMessage
	database.DB.Db.Where("group_id = ?", groupID).Order("created_at ASC").Limit(20).Find(&ChatMessages)

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
	data.CreateChatMessage(*message)
	return c.Status(200).JSON(message)
}

type MessageData struct {
	Message   string `json:"message"`
	Group_id  string `json:"group_id"`
	ChannelID string `json:"channelID"`
}

func HandleCreateChatMessage(c *fiber.Ctx, pusherClient *pusher.Client) error {
	var data MessageData

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	CreateNewChatMessage(c)
	pusherClient.Trigger(data.ChannelID, "new-message", data)

	return c.Status(200).JSON(data)
}
