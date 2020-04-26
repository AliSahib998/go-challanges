package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required,gte=5,lte=10"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token"`
}
