package handlers

import (
	"authentication-go-fiber/model"
	"authentication-go-fiber/security"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var req model.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	var user model.User
	res := db.Where("email = ?", req.Email).First(&user)
	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Error to get the user with the given email",
		})
	}
	if !security.VerifyPassword(user.Password, req.Password) {
		return c.Status(403).JSON(fiber.Map{
			"message": "Incorrect password or email address",
		})
	}
	token, err := security.GenerateToken(user.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(&fiber.Map{
		"token": token,
	})
}
