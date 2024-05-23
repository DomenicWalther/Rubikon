package data

import (
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
)

func CreateChatMessage(message models.GroupChatMessage) error {
	return database.DB.Db.Create(&message).Error
}
