package data

import (
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
)

func GetAllGroups() []models.Group {
	groups := []models.Group{}
	database.DB.Db.Find(&groups)
	return groups
}

func CreateNewGroup(group *models.Group) {
	database.DB.Db.Create(group)
}

func UpdateGroupUserCount(group *models.Group, count int) {
	database.DB.Db.Model(&group).Update("user_count", group.UserCount+count)
}

func UserJoinGroup(user *models.User, group *models.Group) {
	database.DB.Db.Model(&user).Association("Groups").Append(group)
}

func UserLeaveGroup(user *models.User, group *models.Group) {
	database.DB.Db.Model(&user).Association("Groups").Delete(group)
}
