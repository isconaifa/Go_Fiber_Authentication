package config

import (
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	db, err = InitDatabase()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetMysql() *gorm.DB {
	return db
}
