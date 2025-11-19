package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            uint   `gorm:"primary_key"`
	Nome          string `json:"nome"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Idade         int16  `json:"idade"`
	Nacionalidade string `json:"nacionalidade"`
	Genero        string `json:"genero"`
	Curso         string `json:"curso"`
	Graduado      bool   `json:"graduado"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
