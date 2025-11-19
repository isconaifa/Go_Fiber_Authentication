package config

import (
	"authentication-go-fiber/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDatabase() (*gorm.DB, error) {
	dsn := "root:naifa@tcp(localhost:3306)/cadastro_usuarios?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	return db, nil
}
