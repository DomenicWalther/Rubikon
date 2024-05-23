package handlers

import (
	"time"

	"github.com/domenicwalther/rubikon/backend/data"
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pusher/pusher-http-go/v5"
)

type GetChatMessagesResponse struct {
	ID              uuid.UUID `json:"id"`
	GroupID         string    `json:"group_id"`
	Message         string    `json:"message"`
	CreatedAt       time.Time `json:"created_at"`
	Username        string    `json:"username"`
	ProfileImageURL string    `json:"profile_image_url"`
}

func GetChatMessages(c *fiber.Ctx) error {
	groupID := c.Params("id")
	var ChatMessages []models.GroupChatMessage
	// database.DB.Db.Where("group_id = ?", groupID).Order("created_at DESC").Limit(20).Find(&ChatMessages)
	// Get the 20 newest messages, order by created_at with the newest last
	database.DB.Db.Raw("SELECT * FROM group_chat_messages WHERE group_id = ? ORDER BY created_at DESC LIMIT 20", groupID).Scan(&ChatMessages)
	for i, j := 0, len(ChatMessages)-1; i < j; i, j = i+1, j-1 {
		ChatMessages[i], ChatMessages[j] = ChatMessages[j], ChatMessages[i]
	}

	_, clerkData := fetchUserData(getUniqueIDs(ChatMessages))

	// Create new Array out of userData and clerkData
	var ChatMessagesResponse []GetChatMessagesResponse
	for _, message := range ChatMessages {
		for _, clerk := range clerkData {
			if message.UserID == clerk.ID {
				ChatMessagesResponse = append(ChatMessagesResponse, GetChatMessagesResponse{
					ID:              message.ID,
					GroupID:         message.GroupID.String(),
					Message:         message.Message,
					CreatedAt:       message.CreatedAt,
					Username:        clerk.Username,
					ProfileImageURL: clerk.ProfileImageURL,
				})
			}
		}
	}

	return c.Status(200).JSON(ChatMessagesResponse)
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
	Message         string `json:"message"`
	Group_id        string `json:"group_id"`
	ChannelID       string `json:"channelID"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
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

func getUniqueIDs(messages []models.GroupChatMessage) []string {
	uniqueIDs := make(map[string]bool)
	var uniqueIDsList []string
	for _, message := range messages {
		if _, value := uniqueIDs[message.UserID]; !value {
			uniqueIDs[message.UserID] = true
			uniqueIDsList = append(uniqueIDsList, message.UserID)
		}
	}
	return uniqueIDsList
}
