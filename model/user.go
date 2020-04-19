package model

type User struct {
	Username string `sql:"username,pk" json:"username"`
	Password string `sql:"password" json:"password"`
	Token    string `sql:"token" json:"token"`
}
