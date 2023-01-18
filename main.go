package main

import (
	"fmt"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/problemrecord"
	"golang-crud-rest-api/settings"
	"log"

	//"github.com/rs/cors"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// cors := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{
	// 		http.MethodPost,
	// 		http.MethodGet,
	// 	},
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: false,
	// })

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.MigrateUSER()
	database.MigrateSYSTEM()
	database.MigratePROBLEM()
	database.MigrateLEVEL()
	database.MigrateCONTACT()
	database.MigrateANGENCY()
	database.MigratePROBLEMRECORD()
	database.MigrateSTATUS()
	///database.MigrateUPLOAD()

	router := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})

	RegisterProductRoutesfiber(router)

	// router.Use(cors.New())
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Start the server
	log.Printf("Starting Server on port %s\n", AppConfig.Port)
	log.Fatal(router.Listen(fmt.Sprintf(":%s", AppConfig.Port)))

}

func RegisterProductRoutesfiber(router *fiber.App) {

	
	router.Static("/upload", "./uploads")
	router.Post("/user", settings.CreateUser)
	router.Get("/users", settings.GetUsers)
	router.Get("/users/:id", settings.GetUserById)
	router.Patch("/users/:id", settings.UpdateUser)
	router.Delete("/users/:id", settings.DeleteUser)
	router.Post("/system", settings.CreateSystem)
	router.Get("/systems", settings.GetSystems)
	router.Get("/system/:id", settings.GetSystemById)
	router.Patch("/system/:id", settings.UpdateSystem)
	router.Delete("/system/:id", settings.DeleteSystem)
	router.Post("/problem", settings.CreateProblem)
	router.Get("/problems", settings.GetProblems)
	router.Get("/problem/:id", settings.GetProblemById)
	router.Patch("/problem/:id", settings.UpdateProblem)
	router.Delete("/problem/:id", settings.DeleteProblem)
	router.Post("/level", settings.CreateLevel)
	router.Get("/levels", settings.GetLevels)
	router.Get("/level/:id", settings.GetLevelById)
	router.Patch("/level/:id", settings.UpdateLevel)
	router.Delete("/level/:id", settings.DeleteLevel)
	router.Post("/contact", settings.CreateContact)
	router.Get("/contacts", settings.GetContacts)
	router.Get("/contact/:id", settings.GetContactById)
	router.Patch("/contact/:id", settings.UpdateContact)
	router.Delete("/contact/:id", settings.DeleteContact)
	router.Post("/agency", settings.CreateAgency)
	router.Get("/agencys", settings.GetAgencys)
	router.Get("/agency/:id", settings.GetAgencyById)
	router.Patch("/agency/:id", settings.UpdateAgency)
	router.Delete("/agency/:id", settings.DeleteAgency)
	router.Post("/status", settings.CreateStatus)
	router.Get("/statuss", settings.GetStatuss)
	router.Get("/status/:id", settings.GetStatusById)
	router.Patch("/status/:id", settings.UpdateStatus)
	router.Delete("/status/:id", settings.DeleteStatus)

	router.Get("/time/:id", problemrecord.CalculateTime)
	router.Post("/problemrecord", problemrecord.CreateProblemRecord)
	router.Get("/problemrecords", problemrecord.GetProblemRecords)
	router.Get("/problemrecord/:id", problemrecord.GetProblemRecord)
	router.Get("/publiclink/:id", problemrecord.PublicLink)
	router.Patch("/problemupdate/:id", problemrecord.UpdateProblemRecord)
	router.Patch("/problemcompleted/:id", problemrecord.CompletedProblemRecord)
	router.Patch("/problemrecordcancal/:id", problemrecord.CancalProblemRecord)

	router.Get("/problemrecordbyagency/:id", problemrecord.GetProblemRecordByAgency)
	router.Get("/problemrecordbycontact/:id", problemrecord.GetProblemRecordByContact)
	router.Get("/problemrecordbylevel/:id", problemrecord.GetProblemRecordByLevel)
	router.Get("/problemrecordbysystem/:id", problemrecord.GetProblemRecordBySystem)
	router.Get("/problemrecordbyproblem/:id", problemrecord.GetProblemRecordByProblem)
	router.Get("/problemrecordbyimformer/:id", problemrecord.GetProblemRecordByInformer)
	router.Get("/problemrecordbyimformermessage/:id", problemrecord.GetProblemRecordByInformermessage)
	router.Get("/problemrecordbyproblemtype/:id", problemrecord.GetProblemRecordByProblemtype)

}
