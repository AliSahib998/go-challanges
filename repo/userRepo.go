package repo

import (
	"errors"
	"github.com/AliSahib998/go-challanges/model"
	"github.com/jinzhu/gorm"
)

type UserRepo struct{}

type IUserRepo interface {
	CreateUser(user model.User) error
	GetUserByUsername(username string) model.User
	GetUserByUsernameAndPassword(user model.User) bool
}

func (u *UserRepo) CreateUser(user model.User) error {

	Db.Create(&user)

	if user.ID <= 0 {
		return errors.New("failed to create user")
	}

	return nil
}

func (u *UserRepo) GetUserByUsername(username string) model.User {
	user := model.User{}
	Db.Table("users").Where("username=?", username).First(&user)
	return user
}

//optional
func (u *UserRepo) GetUserByUsernameAndPassword(user model.User) bool {

	var res model.User

	var err = Db.Table("users").Where("username=? and password=?", user.Username, user.Password).Find(&res).Error

	if err == gorm.ErrRecordNotFound {
		return false
	}

	return true

}
