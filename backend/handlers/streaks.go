package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type processUserUpdateBody struct {
	UserExperience int `json:"userExperience"`
}

func ProcessUserProgress(c *fiber.Ctx) error {
	userID := c.Locals("sub").(string)
	fmt.Println("UserID: ", userID)
	updateUserExperience(c, userID)

	currentUserStreak, err := getCurrentStreak(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if currentUserStreak == nil {
		currentUserStreak = createFirstStreak(c.Locals("sub").(string))
		return c.Status(200).JSON(currentUserStreak)
	}
	switch lastStreakDate := determineStreakAction(currentUserStreak); lastStreakDate {
	case "Today":
		return c.Status(200).JSON("You already increased your streak today")
	case "Yesterday":
		currentUserStreak = increaseUserStreak(currentUserStreak)

	case "Older":
		err := resetStreak(currentUserStreak)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		currentUserStreak = createNewStreak(currentUserStreak)
	}
	return c.Status(200).JSON(currentUserStreak)
}

func updateUserExperience(c *fiber.Ctx, userID string) error {
	exp := new(processUserUpdateBody)

	if err := c.BodyParser(exp); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	database.DB.Db.Model(&models.User{}).Where("user_id = ?", userID).Update("experience", gorm.Expr("experience + ?", exp.UserExperience))
	return nil
}

func determineStreakAction(streak *models.Streak) string {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println(err)
	}

	now := time.Now().In(loc)
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

func resetStreak(streak *models.Streak) error {
	streak.IsUserCurrentStreak = false
	streak.StreakEnd = &streak.LastStreakDate
	database.DB.Db.Save(&streak)

	return nil
}

func createNewStreak(streak *models.Streak) *models.Streak {
	newStreak := new(models.Streak)
	newStreak.UserID = streak.UserID
	newStreak.StreakStart = time.Now()
	newStreak.IsUserCurrentStreak = true
	newStreak.PreviousStreakID = &streak.ID
	database.DB.Db.Create(&newStreak)

	return newStreak
}

func increaseUserStreak(streak *models.Streak) *models.Streak {
	streak.StreakLength++
	streak.LastStreakDate = time.Now()
	database.DB.Db.Save(&streak)
	return streak
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

func createFirstStreak(userID string) *models.Streak {
	streak := new(models.Streak)
	streak.UserID = userID
	database.DB.Db.Create(&streak)

	return streak
}
