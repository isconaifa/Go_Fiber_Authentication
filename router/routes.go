package router

import (
	"authentication-go-fiber/handlers"
	"authentication-go-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(router *fiber.App) {
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/user", middleware.JWTProtected, handlers.CreateUser)
	v1.Get("/user/:id", handlers.GetUserById)
	v1.Get("/users", handlers.GetAllUsers)
	v1.Put("/user/:id", middleware.JWTProtected, handlers.UpdateUser)
	v1.Delete("/user/:id", middleware.JWTProtected, handlers.DeleteUser)
	v1.Post("/login", handlers.LoginHandler)

	//auth

}
