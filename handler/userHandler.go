package handler

import (
	"github.com/AliSahib998/go-challanges/model"
	"github.com/AliSahib998/go-challanges/repo"
	"github.com/AliSahib998/go-challanges/service"
	"github.com/AliSahib998/go-challanges/util"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type userHandler struct {
	userService service.IUserService
}

const rootPath = "/user"

func NewHandler(router *chi.Mux) *chi.Mux {
	h := &userHandler{
		userService: &service.UserService{UserRepo: &repo.UserRepo{}},
	}
	router.Post(rootPath, h.createUser)

	return router
}

func (u *userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	log.Info("start createUser")
	var user model.User
	err := util.DecodeJson(r.Body, &user)

	if err.Code > 0 {
		util.ErrorHandler(w, err)
	}

	err = u.userService.CreateUser(user)

	if err.Code > 0 {
		util.ErrorHandler(w, err)
	}
	w.Header().Add("Content-Type", "application/json")

	log.Info("end createUser")
}
