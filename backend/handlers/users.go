package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	type Response struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	var response Response
	user := new(models.User)
	if err := c.BodyParser(&response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user.User_ID = response.Data.ID
	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	result := map[string]interface{}{}

	database.DB.Db.Model(&models.User{}).First(&result)
	return c.Status(200).JSON(&result)
}

func GetTopUsers(c *fiber.Ctx) error {
	var topUsers []models.User
	database.DB.Db.Order("experience desc").Limit(10).Find(&topUsers)

	userIds := make([]string, 0, len(topUsers))
	for _, user := range topUsers {
		userIds = append(userIds, user.User_ID)
	}

	jsonData, _ := json.Marshal(userIds)
	fmt.Println(string(jsonData))
	fetchTopUserData(userIds)
	return c.Status(200).JSON(&topUsers)

}

func formatUserIDs(userIDs []string) string {
	var formattedIDs []string
	for _, id := range userIDs {
		formattedIDs = append(formattedIDs, fmt.Sprintf("user_id=%s", id)) // Individual "user_idArray="
	}
	return strings.Join(formattedIDs, "&")
}

func appendUserIDsToURL(baseURL string, userIDs []string) (string, error) {
	formattedIDString := formatUserIDs(userIDs)

	// Construct the query string.
	newURL := fmt.Sprintf("%s?%s", baseURL, formattedIDString)
	return newURL, nil
}

func fetchTopUserData(userIds []string) error {
	//agent := fiber.Get("https://api.clerk.com/v1/users")
	// agent.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CLERK_SECRET")))
	baseURL := "https://api.clerk.com/v1/users"
	newURL, err := appendUserIDsToURL(baseURL, userIds)
	if err != nil {
		return err
	}
	fmt.Println(newURL)
	return nil
}
