package problemrecord

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"
	"github.com/gorilla/mux"
)

//use form-data

func CreateProblemRecord(w http.ResponseWriter, r *http.Request) {
	var problemrecord entities.ProblemRecord
	_ = json.NewDecoder(r.Body).Decode(&problemrecord)
	database.Instance.Create(&problemrecord)
	json.NewEncoder(w).Encode(problemrecord)
	Uploadfile(w, r)
}

func GetProblemRecords(w http.ResponseWriter, r *http.Request) {
	var problemrecords []entities.ProblemRecord
	database.Instance.Find(&problemrecords)
	json.NewEncoder(w).Encode(problemrecords)
}

func GetProblemRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var problemrecord entities.ProblemRecord
	database.Instance.First(&problemrecord, params["id"])
	json.NewEncoder(w).Encode(problemrecord)
}








