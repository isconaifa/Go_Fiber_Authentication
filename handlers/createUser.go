package handlers

import (
	"authentication-go-fiber/model"
	"authentication-go-fiber/security"
	"github.com/gofiber/fiber/v2"
)

// create a new user
func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	hash, err := security.GenerateHashPassword(user.Password)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).SendString("User Created Successfully")
}
