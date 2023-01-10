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

func CreateSystem(c *fiber.Ctx) error {
	var product entities.System
	c.BodyParser(&product)
	database.Instance.Create(&product)
	return c.JSON(product)
}

func GetSystemById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfSystemExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.System
	database.Instance.First(&product, productId)
	return c.JSON(product)
}

func GetSystems(c *fiber.Ctx) error {
	var products []entities.System
	database.Instance.Find(&products)
	return c.JSON(products)
}

func UpdateSystem(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfSystemExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.System
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteSystem(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfSystemExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.System
	database.Instance.First(&product, productId)
	database.Instance.Delete(&product)
	return c.JSON("Product Deleted!")
}

func checkIfSystemExists(id string) bool {
	var product entities.System
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}


// func CreateSystem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product entities.System
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Create(&product)
// 	json.NewEncoder(w).Encode(product)
// }

// func GetSystemById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfSystemExists(productId) == false { 
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.System
// 	database.Instance.First(&product, productId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func GetSystems(w http.ResponseWriter, r *http.Request) {
// 	var products []entities.System
// 	database.Instance.Find(&products)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(products)
// }

// func UpdateSystem(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfSystemExists(productId) == false {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.System
// 	database.Instance.First(&product, productId)
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Save(&product)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func DeleteSystem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	productId := mux.Vars(r)["id"]
// 	if checkIfSystemExists(productId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.System
// 	database.Instance.Delete(&product, productId)
// 	json.NewEncoder(w).Encode("Product Deleted Successfully!")
// }

// func checkIfSystemExists(productId string) bool {
// 	var product entities.System
// 	database.Instance.First(&product, productId)
// 	if product.ID == 0 {
// 		return false
// 	}
// 	return true
// }
