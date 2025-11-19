package handlers

import (
	"authentication-go-fiber/model"
	"github.com/gofiber/fiber/v2"
)

// get all users
func GetAllUsers(c *fiber.Ctx) error {
	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(&users)
}
