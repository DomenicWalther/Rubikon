package handlers

import (
	"github.com/domenicwalther/rubikon/backend/data"
	"github.com/gofiber/fiber/v2"
)

func HandleGetShopSkins(c *fiber.Ctx) error {
	skins := data.GetAllSkins()
	return c.Status(200).JSON(skins)
}
