package models

import (
  "github.com/jinzhu/gorm"
  _"github.com/jinzhu/gorm/dialects/sqlite"
)


var DB *gorm.DB


func ConnectToDatabase() {
	database, err := gorm.Open("sqlite3", "book.db")

	if err != nil {
		panic("Faild conenct to database")
	}

	database.AutoMigrate(&Book{})

	DB = database
}