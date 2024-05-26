package handlers

import (
	"github.com/domenicwalther/rubikon/backend/data"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleGetShopSkins(c *fiber.Ctx) error {
	skins := data.GetAllSkins()
	userSkins := data.GetAllUserSkins(c.Params("id"))

	//remove userSkins from skins
	for i := 0; i < len(skins); i++ {
		for j := 0; j < len(userSkins); j++ {
			if skins[i].ID == userSkins[j].ID {
				skins = append(skins[:i], skins[i+1:]...)
				i--
				break
			}
		}
	}

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
	currentUserLevel := calculateUserLevel(user.Experience)
	if currentUserLevel < skin.RequiredLevel {
		return c.Status(400).JSON(fiber.Map{
			"message": "User level too low",
		})
	}

	err = data.UpdateUserCurrency(&user, -skin.Price)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Not enough currency",
		})
	}
	data.UserBuySkin(&user, &skin)

	return c.Status(200).JSON(&user)
}
