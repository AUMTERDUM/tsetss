package problemrecord

import (
	"encoding/json"
	"net/http"

	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
)

func PublicLink(w http.ResponseWriter, r *http.Request) {
	var problemrecords []entities.ProblemRecord
	database.Instance.Find(&problemrecords)
	json.NewEncoder(w).Encode(problemrecords)
	var problemsender []entities.ProblemSender
	database.Instance.Find(&problemsender)
	json.NewEncoder(w).Encode(problemsender)
	var problemcompleted []entities.CompleteRecord
	database.Instance.Find(&problemcompleted)
	json.NewEncoder(w).Encode(problemcompleted)
}
