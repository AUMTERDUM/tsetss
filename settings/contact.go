package settings

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"

	//"net/http"

	//"github.com/gorilla/mux"
	"github.com/gofiber/fiber/v2"
)

// func CreateContact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product entities.Contact
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Create(&product)
// 	json.NewEncoder(w).Encode(product)
// }

// func GetContactById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfContactExists(productId) == false {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Contact
// 	database.Instance.First(&product, productId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func GetContacts(w http.ResponseWriter, r *http.Request) {
// 	var products []entities.Contact
// 	database.Instance.Find(&products)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(products)
// }

// func UpdateContact(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if checkIfContactExists(productId) == false {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Contact
// 	database.Instance.First(&product, productId)
// 	json.NewDecoder(r.Body).Decode(&product)
// 	database.Instance.Save(&product)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func DeleteContact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	productId := mux.Vars(r)["id"]
// 	if checkIfContactExists(productId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entities.Contact
// 	database.Instance.Delete(&product, productId)
// 	json.NewEncoder(w).Encode("Product Deleted Successfully!")
// }

// func checkIfContactExists(productId string) bool {
// 	var product entities.Contact
// 	database.Instance.First(&product, productId)
// 	if product.ID == 0 {
// 		return false
// 	}
// 	return true
// }

//convert to fiber

func CreateContact(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var product entities.Contact
	c.BodyParser(&product)
	database.Instance.Create(&product)
	c.JSON(product)

	return nil
}

func GetContactById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfContactExists(productId) == false {
		c.JSON("Product Not Found!")
		return nil
	}
	var product entities.Contact
	database.Instance.First(&product, productId)
	c.Set("Content-Type", "application/json")
	c.JSON(product)

	return nil
}	


func GetContacts(c *fiber.Ctx) error {
	var products []entities.Contact
	database.Instance.Find(&products)
	c.Set("Content-Type", "application/json")
	c.Status(200)
	c.JSON(products)

	return nil
}

func UpdateContact(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfContactExists(productId) == false {
		c.JSON("Product Not Found!")
		//return nil
	}
	var product entities.Contact
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	c.Set("Content-Type", "application/json")
	c.JSON(product)

	return nil
}

func DeleteContact(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	productId := c.Params("id")
	if checkIfContactExists(productId) == false {
		c.Status(404)
		c.JSON("Product Not Found!")
	}
	var product entities.Contact
	database.Instance.Delete(&product, productId)
	c.JSON("Product Deleted Successfully!")

	return nil
}

func checkIfContactExists(productId string) bool {
	var product entities.Contact
	database.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}
