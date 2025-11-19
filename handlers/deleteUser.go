package handlers

import (
	"authentication-go-fiber/model"
	"github.com/gofiber/fiber/v2"
)

// delete a new user
func DeleteUser(c *fiber.Ctx) error {
	var user model.User
	id := c.Params("id")
	if id == "" {
		return c.Status(404).SendString("User Not Found")
	}
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(204).SendString("User Deleted")
}
