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

func CreateGroup(c *fiber.Ctx) error {
	group := new(models.Group)
	if err := c.BodyParser(group); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}
	group.OwnerID = c.Locals("sub").(string)
	database.DB.Db.Create(&group)
	return c.Status(200).JSON(group)
}

func JoinGroup(c *fiber.Ctx) error {
	user := models.User{}
	id := new(struct {
		ID int `json:"group_id"`
	})
	err := c.BodyParser(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	database.DB.Db.Where("user_id = ?", c.Locals("sub").(string)).First(&user)

	group := models.Group{ID: uint(id.ID)}
	database.DB.Db.Model(&user).Association("Groups").Append(&group)

	return c.Status(200).JSON(&user)
}
