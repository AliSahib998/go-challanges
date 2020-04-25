package service

import (
	"fmt"
	"github.com/AliSahib998/go-challanges/model"
	"github.com/AliSahib998/go-challanges/repo"
	"github.com/AliSahib998/go-challanges/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repo.IUserRepo
}

type IUserService interface {
	CreateUser(user model.User) util.ErrorResponse
}

func (u *UserService) CreateUser(user model.User) util.ErrorResponse {
	log.Debug("check username to exist")
	fmt.Println(user.Username, user.Password, user.Token)
	var res = u.UserRepo.GetUser(user.Username)

	if res.ID != 0 {
		return util.ErrorResponse{Code: 403, Message: "username is exist"}
	}

	//encrypt password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	//create user
	err := u.UserRepo.CreateUser(user)

	if err != nil {
		log.Debug("error occured", err.Error())
		return util.ErrorResponse{Code: 500, Message: "error occured"}
	}

	return util.ErrorResponse{}
}
