package handlers

import (
	"errors"

	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ManageDailyStreak(c *fiber.Ctx) error {
	currentUserStreak, err := getCurrentStreak(c.Locals("sub").(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if currentUserStreak == nil {
		createFirstStreak(c.Locals("sub").(string))
	}
	return c.Status(200).JSON(currentUserStreak)
}

func getCurrentStreak(userID string) (*models.Streak, error) {
	var currentStreak models.Streak

	result := database.DB.Db.Where(&models.Streak{UserID: userID, IsUserCurrentStreak: true}).First(&currentStreak)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}

	return &currentStreak, nil
}

func createFirstStreak(userID string) error {
	streak := new(models.Streak)
	streak.UserID = userID
	database.DB.Db.Create(&streak)

	return nil
}
