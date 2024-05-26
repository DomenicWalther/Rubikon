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
