package database

import (
	"fmt"
	"log"

	"github.com/FarrellOkello/Go_course/FarrellOkello/Go_course/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "sqlserver://sa:12345678=j@localhost:1433?database=GO_TEST"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	// Create connection pool
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	DB = db
	fmt.Printf("Connected!\n")
	db.AutoMigrate(&models.User{})

}
