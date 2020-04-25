package main

import (
	"github.com/AliSahib998/go-challanges/util"
)

func main() {
	//util.InitLogger()
	//util.InitEnvVars()
	//
	//log.Debug("connect to Db")
	//repo.DbConnect()

	//router := chi.NewRouter()
	//router.Use(middleware.Recoverer)
	//
	//handler.NewHandler(router)
	//handler.NewHealhtHandler(router)
	//
	//port := os.Getenv("Port")
	//log.Debug("Starting server at port:",port)
	//log.Fatal(http.ListenAndServe(":"+port, router))

	type User struct {
		Username string `validate:"required"`
		Tagline  string `validate:"required,lt=10"`
		Tagline2 string `validate:"required,gt=1"`
	}

	user := User{
		Username: "",
		Tagline:  "This tagline is way too long.",
		Tagline2: "1",
	}

	util.Test(user)
	//
	//var a = util.Validation(user)
	//fmt.Println(a)

	//util.Test("sss")
}
