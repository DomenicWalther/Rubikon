package handlers

import (
	"github.com/domenicwalther/rubikon/backend/data"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleGetShopSkins(c *fiber.Ctx) error {
	skins := data.GetAllSkins()
	return c.Status(200).JSON(skins)
}

func HandleBuySkin(c *fiber.Ctx) error {
	id := new(struct {
		ID uuid.UUID `json:"skinID"`
	})
	err := c.BodyParser(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}
	userID := c.Locals("sub").(string)
	user := data.GetUserById(userID)

	skin := data.GetSkinByID(id.ID.String())
	data.UserBuySkin(&user, &skin)

	return c.Status(200).JSON(&user)
}
