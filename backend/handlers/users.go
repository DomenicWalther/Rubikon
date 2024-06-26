package handlers

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/domenicwalther/rubikon/backend/data"
	"github.com/domenicwalther/rubikon/backend/database"
	"github.com/domenicwalther/rubikon/backend/models"
	"github.com/go-jose/go-jose/v3/json"
	"github.com/gofiber/fiber/v2"
)

type LeaderBoardUserInfo struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	FirstName         string `json:"first_name"`
	ProfileImageURL   string `json:"profile_image_url"`
	Experience        int    `json:"experience"`
	MonthlyExperience int    `json:"monthly_experience"`
}

type UserResponse struct {
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
	models.User
}

func HandleCreateUser(c *fiber.Ctx) error {

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

func HandleGetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	_, userData := fetchUserData([]string{id})
	user := data.GetUserById(id)

	userResponse := UserResponse{
		Username:     userData[0].Username,
		ProfileImage: userData[0].ProfileImageURL,
		User:         user,
	}

	return c.Status(200).JSON(&userResponse)
}

func HandleGetTopUsers(c *fiber.Ctx) error {

	topUsers := data.GetUsersWithHighestExperience(10)
	userIds := make([]string, 0, len(topUsers))
	for _, user := range topUsers {
		userIds = append(userIds, user.User_ID)
	}
	userMap := make(map[string]models.User)
	for _, user := range topUsers {
		userMap[user.User_ID] = user
	}

	_, clerkData := fetchUserData(userIds)

	for i := range clerkData {
		userID := clerkData[i].ID
		if matchingUser, exists := userMap[userID]; exists {
			clerkData[i].Experience = matchingUser.Experience
			clerkData[i].MonthlyExperience = matchingUser.MonthlyExperience

		}
	}

	return c.Status(200).JSON(clerkData)

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

func fetchUserData(userIds []string) (error, []LeaderBoardUserInfo) {
	baseURL := "https://api.clerk.com/v1/users"
	newURL, err := appendUserIDsToURL(baseURL, userIds)
	if err != nil {
		return err, nil
	}
	method := "GET"

	payload := strings.NewReader("")
	client := &http.Client{}
	req, err := http.NewRequest(method, newURL, payload)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("CLERK_API"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	var users []LeaderBoardUserInfo

	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	fmt.Println(users)
	return nil, users
}

func calculateUserLevel(experience int) int {
	level := math.Floor(math.Sqrt(float64(experience)/10) + 1)
	return int(level)
}
