package main

import (
	"authentication-go-fiber/config"
	"authentication-go-fiber/handlers"
	"authentication-go-fiber/router"
)

// https://medium.com/code-beyond/go-fiber-jwt-auth-eab51a7e2129
func main() {
	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}
	handlers.SetDB(db)
	err = config.Init()
	if err != nil {
		return
	}
	router.SetupRouter()
}
