package database

import (
	"golang-crud-rest-api/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func MigrateUSER() {
	Instance.AutoMigrate(&entities.User{})
	log.Println("Database Migration Completed...")
}

func MigrateSYSTEM() {
	Instance.AutoMigrate(&entities.System{})
	log.Println("Database Migration Completed...")
}

func MigratePROBLEM() {
	Instance.AutoMigrate(&entities.Problem{})
	log.Println("Database Migration Completed...")
}

func MigrateLEVEL() {
	Instance.AutoMigrate(&entities.Level{})
	log.Println("Database Migration Completed...")
}


