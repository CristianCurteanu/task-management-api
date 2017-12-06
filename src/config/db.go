package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DatabaseConnection() *gorm.DB {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open("sqlite3", dir+"/db/gorm.db")
	if err != nil {
		panic(err.Error())
	} else {
		return db
	}
}
