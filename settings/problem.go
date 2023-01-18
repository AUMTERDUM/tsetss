package settings

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	//"net/http"

	//"github.com/gorilla/mux"
	"github.com/gofiber/fiber/v2"
)

//convert to fiber

func CreateProblem(c *fiber.Ctx) error {
	var product entities.Problemtype
	c.BodyParser(&product)
	database.Instance.Create(&product)
	return c.JSON(product)
}

func GetProblemById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfProblemExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Problemtype
	database.Instance.First(&product, productId)
	return c.JSON(product)
}

func GetProblems(c *fiber.Ctx) error {
	var products []entities.Problemtype
	database.Instance.Find(&products)
	return c.JSON(products)
}

func UpdateProblem(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfProblemExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Problemtype 
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteProblem(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfProblemExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Problemtype
	database.Instance.First(&product, productId)
	database.Instance.Delete(&product)
	return c.JSON("Product Deleted!")
}

func checkIfProblemExists(id string) bool {
	var product entities.Problemtype
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}




// func CreateProblem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product entities.Problem
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Create(&product)
// 	json.NewEncoder(w).Encode(product)
// }

// func GetProblemById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfProblemExists(productId) == false { 
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Problem
// 	database.Instance.First(&product, productId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func GetProblems(w http.ResponseWriter, r *http.Request) {
// 	var products []entities.Problem
// 	database.Instance.Find(&products)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(products)
// }

// func UpdateProblem(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfProblemExists(productId) == false {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Problem
// 	database.Instance.First(&product, productId)
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Save(&product)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func DeleteProblem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	productId := mux.Vars(r)["id"]
// 	if checkIfProblemExists(productId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Problem
// 	database.Instance.Delete(&product, productId)
// 	json.NewEncoder(w).Encode("Product Deleted Successfully!")
// }

// func checkIfProblemExists(productId string) bool {
// 	var product entities.Problem
// 	database.Instance.First(&product, productId)
// 	if product.ID == 0 {
// 		return false
// 	}
// 	return true
// }
