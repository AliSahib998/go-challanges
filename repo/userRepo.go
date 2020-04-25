package repo

import (
	"errors"
	"github.com/AliSahib998/go-challanges/model"
)

type UserRepo struct{}

type IUserRepo interface {
	CreateUser(user model.User) error
	GetUser(username string) model.User
}

func (u *UserRepo) CreateUser(user model.User) error {

	Db.Create(&user)

	if user.ID <= 0 {
		return errors.New("failed to create user")
	}

	return nil
}

func (u *UserRepo) GetUser(username string) model.User {
	user := model.User{}
	Db.Table("users").Where("username=?", username).First(&user)
	return user
}
