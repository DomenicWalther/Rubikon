package handlers

import (
	"errors"
	"time"

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

func checkStreakDate(streak *models.Streak) string {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	yesterday := today.AddDate(0, 0, -1)

	if streak.LastStreakDate.After(yesterday) && streak.LastStreakDate.Before(today) {
		return "Yesterday"
	} else if streak.LastStreakDate.Equal(today) || streak.LastStreakDate.After(today) {
		return "Today"
	} else {
		return "Older"
	}
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
