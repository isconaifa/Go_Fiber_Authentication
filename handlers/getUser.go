package handlers

import (
	"authentication-go-fiber/model"
	"github.com/gofiber/fiber/v2"
)

// find user by id
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(404).SendString("User Not Found")
	}
	var user model.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.Status(200).JSON(&user)
}
