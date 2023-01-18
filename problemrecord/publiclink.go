package problemrecord

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	//"net/http"
	"github.com/gofiber/fiber/v2"

)

func PublicLink(c *fiber.Ctx) error {
	id := c.Params("id")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("id = ?",id).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	
	return c.JSON(problemrecord)
}