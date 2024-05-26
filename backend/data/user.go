package data

import (
	"errors"

	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
)

func GetUserById(userID string) models.User {
	user := models.User{}
	database.DB.Db.Where("user_id = ?", userID).First(&user)
	return user
}

func GetUsersWithHighestExperience(limit int) []models.User {
	topUsers := []models.User{}
	database.DB.Db.Order("experience desc").Limit(limit).Find(&topUsers)

	return topUsers
}

func UpdateUserCurrency(user *models.User, amount int) error {
	if user.Currency+amount < 0 {
		return errors.New("not enough currency")
	}
	user.Currency += amount
	database.DB.Db.Save(user)
	return nil
}
