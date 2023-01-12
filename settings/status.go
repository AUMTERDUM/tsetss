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
func CreateStatus(c *fiber.Ctx) error {
	var product entities.Status
	c.BodyParser(&product)
	database.Instance.Create(&product)
	return c.JSON(product)
}

func GetStatusById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfStatusExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Status
	database.Instance.First(&product, productId)
	return c.JSON(product)
}

func GetStatuss(c *fiber.Ctx) error {
	var products []entities.Status
	database.Instance.Find(&products)
	return c.JSON(products)
}

func UpdateStatus(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfStatusExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Status
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteStatus(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfStatusExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Status
	database.Instance.First(&product, productId)
	database.Instance.Delete(&product)
	return c.JSON("Product Deleted!")
}

func checkIfStatusExists(id string) bool {
	var product entities.Status
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}



// func CreateStatus(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product entities.Status
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Create(&product)
// 	json.NewEncoder(w).Encode(product)
// }

// func GetStatusById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfStatusExists(productId) == false { 
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Status
// 	database.Instance.First(&product, productId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func GetStatuss(w http.ResponseWriter, r *http.Request) {
// 	var products []entities.Status
// 	database.Instance.Find(&products)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(products)
// }

// func UpdateStatus(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfStatusExists(productId) == false {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Status
// 	database.Instance.First(&product, productId)
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Save(&product)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func DeleteStatus(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	productId := mux.Vars(r)["id"]
// 	if checkIfStatusExists(productId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Status
// 	database.Instance.Delete(&product, productId)
// 	json.NewEncoder(w).Encode("Product Deleted Successfully!")
// }

// func checkIfStatusExists(productId string) bool {
// 	var product entities.Status
// 	database.Instance.First(&product, productId)
// 	if product.ID == 0 {
// 		return false
// 	}
// 	return true
// }
