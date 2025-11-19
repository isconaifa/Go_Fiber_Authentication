package handlers

import (
	"authentication-go-fiber/model"
	"authentication-go-fiber/security"
	"github.com/gofiber/fiber/v2"
)

// update user
func UpdateUser(c *fiber.Ctx) error {
	user := new(model.User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	hash, err := security.GenerateHashPassword(user.Password)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	user.Password = hash
	if err := db.Where("id = ?", id).Updates(&user).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(&user)
}
