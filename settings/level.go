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
func CreateLevel(c *fiber.Ctx) error {
	var product entities.Level
	c.BodyParser(&product)
	database.Instance.Create(&product)
	return c.JSON(product)
}

func GetLevelById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfLevelExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Level
	database.Instance.First(&product, productId)
	return c.JSON(product)
}

func GetLevels(c *fiber.Ctx) error {
	var products []entities.Level
	database.Instance.Find(&products)
	return c.JSON(products)
}

func UpdateLevel(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfLevelExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Level
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteLevel(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfLevelExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Level
	database.Instance.First(&product, productId)
	database.Instance.Delete(&product)
	return c.JSON("Product Deleted!")
}

func checkIfLevelExists(id string) bool {
	var product entities.Level
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}



// func CreateLevel(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product entities.Level
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Create(&product)
// 	json.NewEncoder(w).Encode(product)
// }

// func GetLevelById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfLevelExists(productId) == false { 
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Level
// 	database.Instance.First(&product, productId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func GetLevels(w http.ResponseWriter, r *http.Request) {
// 	var products []entities.Level
// 	database.Instance.Find(&products)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(products)
// }

// func UpdateLevel(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfLevelExists(productId) == false {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Level
// 	database.Instance.First(&product, productId)
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Save(&product)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func DeleteLevel(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	productId := mux.Vars(r)["id"]
// 	if checkIfLevelExists(productId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Level
// 	database.Instance.Delete(&product, productId)
// 	json.NewEncoder(w).Encode("Product Deleted Successfully!")
// }

// func checkIfLevelExists(productId string) bool {
// 	var product entities.Level
// 	database.Instance.First(&product, productId)
// 	if product.ID == 0 {
// 		return false
// 	}
// 	return true
// }
