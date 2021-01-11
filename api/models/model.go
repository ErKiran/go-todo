package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func Init() *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(os.Getenv("DB_TYPE"), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	))

	if err != nil {
		log.Fatal(err, "Couldnot connect to database")
	}

	return db
}
