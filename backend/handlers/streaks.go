package handlers

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type processUserUpdateBody struct {
	UserExperience int `json:"userExperience"`
}

type processUserUpdateResponse struct {
	StreakLength     int `json:"streakLength"`
	GainedExperience int `json:"gainedExperience"`
}

func ProcessUserProgress(c *fiber.Ctx) error {
	userID := c.Locals("sub").(string)
	fmt.Println("UserID: ", userID)

	rubikonLength := new(processUserUpdateBody)

	if err := c.BodyParser(rubikonLength); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	currentUserStreak, err := getCurrentStreak(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if currentUserStreak == nil {
		currentUserStreak = createFirstStreak(c.Locals("sub").(string))
	}
	switch lastStreakDate := determineStreakAction(currentUserStreak); lastStreakDate {

	case "Yesterday":
		currentUserStreak = increaseUserStreak(currentUserStreak)

	case "Older":
		err := resetStreak(currentUserStreak)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		currentUserStreak = createNewStreak(currentUserStreak)
	}
	experienceGain := calculateUserExperience(rubikonLength.UserExperience, currentUserStreak.StreakLength)
	updateUserExperience(experienceGain, userID)

	response := new(processUserUpdateResponse)
	response.StreakLength = currentUserStreak.StreakLength
	response.GainedExperience = experienceGain
	return c.Status(200).JSON(response)
}

func updateUserExperience(experience int, userID string) error {
	database.DB.Db.Model(&models.User{}).Where("user_id = ?", userID).Update("experience", gorm.Expr("experience + ?", experience))
	database.DB.Db.Model(&models.User{}).Where("user_id = ?", userID).Update("monthly_experience", gorm.Expr("monthly_experience + ?", experience))
	return nil
}

func calculateUserExperience(experience, streakLength int) int {
	experience = experience / 60
	gain := int(math.Ceil(float64(experience) * math.Log(math.Pow(float64(streakLength), 1.2))))
	if gain > experience {
		return gain
	}
	return experience
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
