package main

import (
	"github.com/AliSahib998/go-challanges/handler"
	mid "github.com/AliSahib998/go-challanges/middleware"
	"github.com/AliSahib998/go-challanges/repo"
	"github.com/AliSahib998/go-challanges/util"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	util.InitLogger()
	util.InitEnvVars()

	log.Debug("connect to Db")
	repo.DbConnect()

	router := chi.NewRouter()
	router.Use(mid.JwtAuthentication)

	handler.NewHandler(router)
	handler.NewHealthHandler(router)
	handler.NewCovidHandler(router)

	port := os.Getenv("Port")
	log.Debug("Starting server at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
