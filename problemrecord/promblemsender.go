package problemrecord

import (
	"encoding/json"
	"net/http"

	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
)

func SenderProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.ProblemSender
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}
