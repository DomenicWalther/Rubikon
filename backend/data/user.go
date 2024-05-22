package data

import (
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

}
