package repo

import (
	"fmt"
	"github.com/AliSahib998/go-challanges/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var Db *gorm.DB

func DbConnect() {

	username := os.Getenv("Db_User")
	password := os.Getenv("Db_Pass")
	dbName := os.Getenv("Db_Name")
	dbHost := os.Getenv("Db_Host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	conn.AutoMigrate(&model.User{})

	Db = conn
}
