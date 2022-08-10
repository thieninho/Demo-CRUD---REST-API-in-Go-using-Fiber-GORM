package database

import (
	"log"
	"os"

	"github.com/sixfwa/fiber-gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {

		log.Fatal("Failed to connect \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running")

	//TODO
	db.AutoMigrate(&models.Page{})

	Database = DbInstance{Db: db}

}
