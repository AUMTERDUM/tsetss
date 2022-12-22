package main

import (
	"fmt"
	"golang-crud-rest-api/settings"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"github.com/rs/cors"
)

var DB *gorm.DB

func main() {

    cors := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{
            http.MethodPost,
            http.MethodGet,
        },
        AllowedHeaders:   []string{"*"},
        AllowCredentials: false,
    })
	
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


	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	RegisterProductRoutes(router)
	handler := cors.Handler(router)


	//test:= fmt.Sprintf("Starting Server on port %s", AppConfig.Port)
	// Start the server
	log.Printf("Starting Server on port %s\n", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", AppConfig.Port), handler))
	//http.ListenAndServe(":8080", handler)
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/user", settings.CreateUser).Methods("POST")
	router.HandleFunc("/users", settings.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", settings.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", settings.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", settings.DeleteUser).Methods("DELETE")
	router.HandleFunc("/system", settings.CreateSystem).Methods("POST")
	router.HandleFunc("/systems", settings.GetSystems).Methods("GET")
	router.HandleFunc("/system/{id}", settings.GetSystemById).Methods("GET")
	router.HandleFunc("/system/{id}", settings.UpdateSystem).Methods("PATCH")
	router.HandleFunc("/system/{id}", settings.DeleteSystem).Methods("DELETE")
	router.HandleFunc("/problem", settings.CreateProblem).Methods("POST")
	router.HandleFunc("/problems", settings.GetProblems).Methods("GET")
	router.HandleFunc("/problem/{id}", settings.GetProblemById).Methods("GET")
	router.HandleFunc("/problem/{id}", settings.UpdateProblem).Methods("PATCH")
	router.HandleFunc("/problem/{id}", settings.DeleteProblem).Methods("DELETE")
	router.HandleFunc("/level", settings.CreateLevel).Methods("POST")
	router.HandleFunc("/levels", settings.GetLevels).Methods("GET")
	router.HandleFunc("/level/{id}", settings.GetLevelById).Methods("GET")
	router.HandleFunc("/level/{id}", settings.UpdateLevel).Methods("PATCH")
	router.HandleFunc("/level/{id}", settings.DeleteLevel).Methods("DELETE")
	router.HandleFunc("/contact", settings.CreateContact).Methods("POST")
	router.HandleFunc("/contacts", settings.GetContacts).Methods("GET")
	router.HandleFunc("/contact/{id}", settings.GetContactById).Methods("GET")
	router.HandleFunc("/contact/{id}", settings.UpdateContact).Methods("PATCH")
	router.HandleFunc("/contact/{id}", settings.DeleteContact).Methods("DELETE")
	router.HandleFunc("/agency", settings.CreateAgency).Methods("POST")
	router.HandleFunc("/agencys", settings.GetAgencys).Methods("GET")
	router.HandleFunc("/agency/{id}", settings.GetAgencyById).Methods("GET")
	router.HandleFunc("/agency/{id}", settings.UpdateAgency).Methods("PATCH")                                         
	router.HandleFunc("/agency/{id}", settings.DeleteAgency).Methods("DELETE")

}

