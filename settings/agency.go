package settings

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"

	//"github.com/gorilla/mux"
	"github.com/gofiber/fiber/v2"
)

// func CreateAgency(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product entities.Agency
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Create(&product)
// 	json.NewEncoder(w).Encode(product)
// }

// func GetAgencyById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfAgencyExists(productId) == false { 
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Agency
// 	database.Instance.First(&product, productId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func GetAgencys(w http.ResponseWriter, r *http.Request) {
// 	var products []entities.Agency
// 	database.Instance.Find(&products)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(products)
// }

// func UpdateAgency(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfAgencyExists(productId) == false {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Agency
// 	database.Instance.First(&product, productId)
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Save(&product)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func DeleteAgency(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	productId := mux.Vars(r)["id"]
// 	if checkIfAgencyExists(productId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Agency
// 	database.Instance.Delete(&product, productId)
// 	json.NewEncoder(w).Encode("Product Deleted Successfully!")
// }

// func checkIfAgencyExists(productId string) bool {
// 	var product entities.Agency
// 	database.Instance.First(&product, productId)
// 	if product.ID == 0 {
// 		return false
// 	}
// 	return true
// }

//convert to fiber

func CreateAgency(c *fiber.Ctx) error {
c.Set("Content-Type", "application/json")
	var product entities.Agency
	c.BodyParser(&product)
	database.Instance.Create(&product)
	c.JSON(product)
	return c.JSON(product)
}

func GetAgencyById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfAgencyExists(productId) == false { 
		c.JSON("Product Not Found!")
		
	}
	var product entities.Agency
	database.Instance.First(&product, productId)
	c.Set("Content-Type", "application/json")
	c.JSON(product)
	return c.JSON(product)
}

func GetAgencys(c *fiber.Ctx) error {
	var products []entities.Agency
	database.Instance.Find(&products)
	c.Set("Content-Type", "application/json")
	c.Status(http.StatusOK)
	c.JSON(products)
	return nil
}

func UpdateAgency(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfAgencyExists(productId) == false {
		c.JSON("Product Not Found!")
		return nil
	}
	var product entities.Agency
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	c.Set("Content-Type", "application/json")
	c.JSON(product)
	return nil
}

func DeleteAgency(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	productId := c.Params("id")
	if checkIfAgencyExists(productId) == false {
		c.Status(http.StatusNotFound)
		c.JSON("Product Not Found!")
		return nil
	}
	var product entities.Agency
	database.Instance.Delete(&product, productId)
	c.JSON("Product Deleted Successfully!")
	return nil
}

func checkIfAgencyExists(productId string) bool {
	var product entities.Agency
	database.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}



