package problemrecord

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	//"net/http"
	"github.com/gofiber/fiber/v2"
)


func PublicLink(c *fiber.Ctx) error {//func PublicLink(w http.ResponseWriter, r *http.Request) {
	var problemrecords []entities.ProblemRecord
	database.Instance.Find(&problemrecords)//.Where("problem_id = ?", problemid)
	c.JSON(problemrecords)
	return nil
	
}	