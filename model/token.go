package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Username string
	jwt.StandardClaims
}

type AuthToken struct {
	Username  string `json:"username"`
	AuthToken string `json:"authToken"`
}
