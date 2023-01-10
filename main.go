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
	database.MigratePROBLEMSENDER()
	database.MigratePROBLEMCOMPLETED()
	///database.MigrateUPLOAD()

	router := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})

	// Initialize the router
	//router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	RegisterProductRoutesfiber(router)
	//handler := cors.Handler(router)

	// router.Use(cors.New())
	router.Use(cors.New())

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// 	}))

	//test:= fmt.Sprintf("Starting Server on port %s", AppConfig.Port)
	// Start the server
	log.Printf("Starting Server on port %s\n", AppConfig.Port)
	log.Fatal(router.Listen(fmt.Sprintf(":%s", AppConfig.Port)))
	//http.ListenAndServe(":8080", handler)

}

// func RegisterProductRoutes(router *mux.Router) {
// 	router.HandleFunc("/user", settings.CreateUser).Methods("POST")
// 	router.HandleFunc("/users", settings.GetUsers).Methods("GET")
// 	router.HandleFunc("/users/{id}", settings.GetUserById).Methods("GET")
// 	router.HandleFunc("/users/{id}", settings.UpdateUser).Methods("PATCH")
// 	router.HandleFunc("/users/{id}", settings.DeleteUser).Methods("DELETE")
// 	router.HandleFunc("/system", settings.CreateSystem).Methods("POST")
// 	router.HandleFunc("/systems", settings.GetSystems).Methods("GET")
// 	router.HandleFunc("/system/{id}", settings.GetSystemById).Methods("GET")
// 	router.HandleFunc("/system/{id}", settings.UpdateSystem).Methods("PATCH")
// 	router.HandleFunc("/system/{id}", settings.DeleteSystem).Methods("DELETE")
// 	router.HandleFunc("/problem", settings.CreateProblem).Methods("POST")
// 	router.HandleFunc("/problems", settings.GetProblems).Methods("GET")
// 	router.HandleFunc("/problem/{id}", settings.GetProblemById).Methods("GET")
// 	router.HandleFunc("/problem/{id}", settings.UpdateProblem).Methods("PATCH")
// 	router.HandleFunc("/problem/{id}", settings.DeleteProblem).Methods("DELETE")
// 	router.HandleFunc("/level", settings.CreateLevel).Methods("POST")
// 	router.HandleFunc("/levels", settings.GetLevels).Methods("GET")
// 	router.HandleFunc("/level/{id}", settings.GetLevelById).Methods("GET")
// 	router.HandleFunc("/level/{id}", settings.UpdateLevel).Methods("PATCH")
// 	router.HandleFunc("/level/{id}", settings.DeleteLevel).Methods("DELETE")
// 	router.HandleFunc("/contact", settings.CreateContact).Methods("POST")
// 	router.HandleFunc("/contacts", settings.GetContacts).Methods("GET")
// 	router.HandleFunc("/contact/{id}", settings.GetContactById).Methods("GET")
// 	router.HandleFunc("/contact/{id}", settings.UpdateContact).Methods("PATCH")
// 	router.HandleFunc("/contact/{id}", settings.DeleteContact).Methods("DELETE")
// 	router.HandleFunc("/agency", settings.CreateAgency).Methods("POST")
// 	router.HandleFunc("/agencys", settings.GetAgencys).Methods("GET")
// 	router.HandleFunc("/agency/{id}", settings.GetAgencyById).Methods("GET")
// 	router.HandleFunc("/agency/{id}", settings.UpdateAgency).Methods("PATCH")
// 	router.HandleFunc("/agency/{id}", settings.DeleteAgency).Methods("DELETE")
// 	router.HandleFunc("/problemrecord", problemrecord.CreateProblemRecord).Methods("POST")
// 	router.HandleFunc("/problemrecords", problemrecord.GetProblemRecords).Methods("GET")
// 	router.HandleFunc("/problemrecord/{id}", problemrecord.GetProblemRecord).Methods("GET")
// 	router.HandleFunc("/problemsender", problemrecord.SenderProblem).Methods("POST")
// 	router.HandleFunc("/problemcompleted", problemrecord.ProblemCompleted).Methods("POST")
// 	router.HandleFunc("/publiclink", problemrecord.PublicLink).Methods("GET")
// 	//router.HandleFunc("/problemrecord/upload", problemrecord.Uploadfile).Methods("POST")

// }

func RegisterProductRoutesfiber(router *fiber.App) {

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
	// router.Post("/contact", settings.CreateContact)
	// router.Get("/contacts", settings.GetContacts)
	// router.Get("/contact/:id", settings.GetContactById)
	// router.Patch("/contact/:id", settings.UpdateContact)
	// router.Delete("/contact/:id", settings.DeleteContact)
	// router.Post("/agency", settings.CreateAgency)
	// router.Get("/agencys", settings.GetAgencys)
	// router.Get("/agency/:id", settings.GetAgencyById)
	// router.Patch("/agency/:id", settings.UpdateAgency)
	// router.Delete("/agency/:id", settings.DeleteAgency)
	router.Post("/problemrecord", problemrecord.CreateProblemRecord)
	// router.Get("/problemrecords", problemrecord.GetProblemRecords)
	// router.Get("/problemrecord/:id", problemrecord.GetProblemRecord)
	// router.Post("/problemsender", problemrecord.SenderProblem)
	// router.Post("/problemcompleted", problemrecord.ProblemCompleted)
	// router.Get("/publiclink", problemrecord.PublicLink)
	//router.Post("/problemrecord/upload", problemrecord.Uploadfile)

}
