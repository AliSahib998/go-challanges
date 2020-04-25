package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required,lte=5`
	Password string `json:"password" validate:"required`
	Token    string `json:"token"`
}

type Test struct {
	Username string `json:"username" validate:"required,lte=5`
	Password string `json:"password" validate:"required`
	Token    string `json:"token"`
}
