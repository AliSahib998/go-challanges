package service

import (
	"fmt"
	"github.com/AliSahib998/go-challanges/model"
	"github.com/AliSahib998/go-challanges/repo"
	"github.com/AliSahib998/go-challanges/util"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type UserService struct {
	UserRepo repo.IUserRepo
}

type IUserService interface {
	CreateUser(user model.User) util.RestErrorResponse
	Login(user model.User) (model.AuthToken, util.RestErrorResponse)
}

func (u *UserService) CreateUser(user model.User) util.RestErrorResponse {

	//validation
	log.Debug("validation check")
	var errorMap = util.Validation(user)

	if len(errorMap) > 0 {
		log.Debug("validation error")
		return util.RestErrorResponse{Code: 400, ValidationError: errorMap}
	}

	//check username
	log.Debug("check username is existing")
	var res = u.UserRepo.GetUserByUsername(user.Username)

	if res.ID != 0 {
		return util.RestErrorResponse{Code: 403, Message: "username is exist"}
	}

	//encrypt password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	//create user
	err := u.UserRepo.CreateUser(user)

	if err != nil {
		log.Debug("error occurred", err.Error())
		return util.RestErrorResponse{Code: 500, Message: "error occurred"}
	}

	return util.RestErrorResponse{}
}

func (u *UserService) Login(user model.User) (model.AuthToken, util.RestErrorResponse) {

	//check username and password
	log.Debug("checking process for sign in")
	var res = u.UserRepo.GetUserByUsername(user.Username)
	err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password))
	fmt.Println(err)

	if err != nil {
		return model.AuthToken{}, util.RestErrorResponse{Code: 404, Message: "username and password are wrong"}
	}

	//jwt token create
	tk := model.Token{Username: res.Username}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("Token_Key")))

	return model.AuthToken{Username: res.Username, AuthToken: tokenString}, util.RestErrorResponse{}
}
