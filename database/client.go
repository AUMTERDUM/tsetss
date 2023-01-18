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
	log.Println("Database USER Completed...")
}

func MigrateSYSTEM() {
	Instance.AutoMigrate(&entities.System{})
	log.Println("Database SYSTEM Completed...")
}

func MigratePROBLEM() {
	Instance.AutoMigrate(&entities.Problemtype{})
	log.Println("Database PROBLEM Completed...")
}

func MigrateLEVEL() {
	Instance.AutoMigrate(&entities.Level{})
	log.Println("Database LEVEL Completed...")
}

func MigrateCONTACT() {
	Instance.AutoMigrate(&entities.Contact{})
	log.Println("Database CONTACT Completed...")
}

func MigrateANGENCY() {
	Instance.AutoMigrate(&entities.Agency{})
	log.Println("Database AGENCY Completed...")
}

func MigratePROBLEMRECORD() {
	Instance.AutoMigrate(&entities.ProblemRecord{})
	log.Println("Database PROBLEMRECORD Completed...")
}

func MigrateSTATUS() {
	Instance.AutoMigrate(&entities.Status{})
	log.Println("Database STATUS Completed...")
}






// func MigrateUPLOAD() {
// 	Instance.AutoMigrate(&entities.File{})
// 	log.Println("Database UPLOAD Completed...")
// }




