package handlers

import (
	"github.com/domenicwalther/rubikon/backend/data"
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type GroupWithStatus struct {
	models.Group
	IsMember bool `json:"is_member"`
}

func HandleGetGroups(c *fiber.Ctx) error {
	userHasGroups := userGroups(c)
	groups := data.GetAllGroups()
	resultGroups := make([]GroupWithStatus, 0, len(groups))

	for _, group := range groups {
		joined := false
		for _, userGroup := range *userHasGroups {
			if group.ID == userGroup.ID {
				joined = true
				break
			}
		}
		resultGroups = append(resultGroups, GroupWithStatus{
			Group:    group,
			IsMember: joined,
		})
	}
	return c.Status(200).JSON(resultGroups)
}

func HandleCreateGroup(c *fiber.Ctx) error {
	group := new(models.Group)
	if err := c.BodyParser(group); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}
	group.OwnerID = c.Locals("sub").(string)
	data.CreateNewGroup(group)
	data.UserJoinGroup(&models.User{User_ID: group.OwnerID}, group)

	return c.Status(200).JSON(group)
}

func JoinGroup(c *fiber.Ctx) error {
	id := new(struct {
		ID uuid.UUID `json:"group_id"`
	})
	err := c.BodyParser(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	userID := c.Locals("sub").(string)
	user := data.GetUserById(userID)
	database.DB.Db.Model(&user).Preload("Groups").Association("Groups").Find(&user.Groups)

	group := models.Group{ID: id.ID}
	data.UserJoinGroup(&user, &group)
	data.UpdateGroupUserCount(&group, 1)

	return c.Status(200).JSON(&user)
}

func userGroups(c *fiber.Ctx) *[]models.Group {
	id := c.Params("id")
	user := models.User{}
	database.DB.Db.Where("user_id = ?", id).Preload("Groups").First(&user)

	return &user.Groups
}

func LeaveGroup(c *fiber.Ctx) error {
	id := new(struct {
		ID uuid.UUID `json:"group_id"`
	})
	err := c.BodyParser(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	user := data.GetUserById(c.Locals("sub").(string))
	group := models.Group{ID: id.ID}
	data.UserLeaveGroup(&user, &group)
	data.UpdateGroupUserCount(&group, -1)

	return c.Status(200).JSON(&user)
}
