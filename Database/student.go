package database

import (
	"fiber_api/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}
var Database DbInstance

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("api.db"),&gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n",err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migration")
    
	db.AutoMigrate(&models.Student{})   // create separate table for each one

	Database = DbInstance{Db: db}
}
