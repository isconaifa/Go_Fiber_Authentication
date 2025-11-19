package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func SetupRouter() {
	router := fiber.New()
	router.Use(logger.New())
	initializeRoutes(router)
	fmt.Println("Server is up and running on port 3000")
	log.Fatal(router.Listen(":3000"))

}
