package handlers

import (
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	type Response struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	var response Response
	user := new(models.User)
	if err := c.BodyParser(&response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user.User_ID = response.Data.ID
	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	result := map[string]interface{}{}

	database.DB.Db.Model(&models.User{}).First(&result)
	return c.Status(200).JSON(&result)
}
