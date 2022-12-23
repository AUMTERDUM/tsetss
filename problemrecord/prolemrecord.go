package problemrecord

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProblemRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.ProblemRecord
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProblemRecordById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProblemRecordExists(productId) == false { 
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.ProblemRecord
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProblemRecords(w http.ResponseWriter, r *http.Request) {
	var products []entities.ProblemRecord
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func checkIfProblemRecordExists(id string) bool {
	var product entities.ProblemRecord
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}
