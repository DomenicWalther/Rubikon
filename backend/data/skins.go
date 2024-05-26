package data

import (
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
)

func GetAllSkins() []models.Skin {
	skins := []models.Skin{}
	database.DB.Db.Find(&skins)
	return skins
}

func GetAllUserSkins(userID string) []models.Skin {
	user := GetUserById(userID)
	skins := []models.Skin{}
	database.DB.Db.Model(&user).Association("Skins").Find(&skins)
	return skins
}

func UserBuySkin(user *models.User, skin *models.Skin) {
	database.DB.Db.Model(&user).Association("Skins").Append(skin)
}

func GetSkinByID(skinID string) models.Skin {
	skin := models.Skin{}
	database.DB.Db.Where("id = ?", skinID).First(&skin)
	return skin
}
