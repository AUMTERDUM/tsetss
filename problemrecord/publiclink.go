package problemrecord

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	//"net/http"
	"github.com/gofiber/fiber/v2"
)

// func PublicLink(w http.ResponseWriter, r *http.Request) {
// 	var problemrecords []entities.ProblemRecord
// 	database.Instance.Find(&problemrecords)
// 	json.NewEncoder(w).Encode(problemrecords)
// 	var problemsender []entities.ProblemSender
// 	database.Instance.Find(&problemsender)
// 	json.NewEncoder(w).Encode(problemsender)
// 	var problemcompleted []entities.CompleteRecord
// 	database.Instance.Find(&problemcompleted)
// 	json.NewEncoder(w).Encode(problemcompleted)
// }

//fiber

func PublicLink(c *fiber.Ctx) error {
	var problemrecords []entities.ProblemRecord
	database.Instance.Find(&problemrecords)
	c.JSON(problemrecords)
	return nil
	

}