package problemrecord

import (
	//"encoding/json"
	"fmt"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"
	"os"
	"path"
	"path/filepath"

	//"golang.org/x/crypto/bcrypt"
	//"github.com/gorilla/mux"
	"github.com/gofiber/fiber/v2"
)

//use form-data

// func CreateProblemRecord(c *fiber.Ctx) {
// 	file, err := c.FormFile("problem_records")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	//get file name
// 	fileName := file.Filename
// 	//get file extension
// 	ext := filepath.Ext(file.Filename)
// 	//get file size
// 	size := file.Size
// 	//get file path
// 	filePath := "/upload/" + fileName + ext
// 	fmt.Println("File Info")
// 	fmt.Println("File Name:", fileName)
// 	fmt.Println("File Size:", size)
// 	fmt.Println("File Path:", filePath)
// 	fmt.Println("File Extension:", ext)
// 	filename := path.Join("uploads", path.Base(file.Filename))
// 	dest, err := os.Create(filename)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	defer dest.Close()
// 	// Copy data in file
// 	if err := c.SaveFile(file, dest.Name()); err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	agency := c.FormValue("agency")
// 	contact := c.FormValue("contact")
// 	problem := c.FormValue("problem")
// 	level := c.FormValue("level")
// 	informer := c.FormValue("informer")
// 	informermessage := c.FormValue("informermessage")
// 	system := c.FormValue("system")
// 	problemtype := c.FormValue("problemtype")

// 	c.Set("Content-Type", "application/json")

// 	problemrecord := entities.ProblemRecord{
// 		File_name:       fileName,
// 		Path_file:       filePath,
// 		Agency:          agency,
// 		Contact:         contact,
// 		Problem:         problem,
// 		Level:           level,
// 		Informer:        informer,
// 		Informermessage: informermessage,
// 		System:          system,
// 		Problemtype:     problemtype,
// 		Status:          1,

// 		File_extension: ext,
// 		File_size:      int(size),
// 	}

// 	database.Instance.Create(&problemrecord)
// 	c.JSON(problemrecord)
// }

func CreateProblemRecord(c *fiber.Ctx) error {
	file, err := c.FormFile("problem_records")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}


	//get file name
	fileName := file.Filename
	//get file extension
	ext := filepath.Ext(file.Filename)
	//get file size
	size := file.Size
	//get file path
	filePath := "/upload/" + fileName + ext
	//log file data
	fmt.Println("File Info")
	fmt.Println("File Name:", fileName)
	fmt.Println("File Size:", size)
	fmt.Println("File Path:", filePath)
	fmt.Println("File Extension:", ext)

	//upload file in folder
	filename := path.Join("./uploads/", path.Base(fileName))
	//create file
	dest, err := os.Create(filename)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer dest.Close()
	// Copy data in file
	if err := c.SaveFile(file, dest.Name()); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// //check service
	// id_service := c.FormValue("id_service")
	// var service models.Service
	// database.Instance.Select("id_service").Where("id_service = ?", id_service).Find(&service)
	// if id_service != service.Id_Service {
	// 	return c.JSON(fiber.Map{
	// 		"massage": "Service not Found",
	// 		"status":  http.StatusBadRequest,
	// 	})
	// }

	// strService := c.FormValue("service_id")
	// Idservice, err := strconv.Atoi(strService)
	// if err != nil {
	// 	panic(err)
	// }
	agency := c.FormValue("agency")
	contact := c.FormValue("contact")
	problem := c.FormValue("problem")
	level := c.FormValue("level")
	informer := c.FormValue("informer")
	informermessage := c.FormValue("informermessage")
	system := c.FormValue("system")
	problemtype := c.FormValue("problemtype")
	
	problemrecord := entities.ProblemRecord{
		File_name:       fileName,
		Path_file:       filePath,
		Agency:          agency,
		Contact:         contact,
		Problem:         problem,
		Level:           level,
		Informer:        informer,
		Informermessage: informermessage,
		System:          system,
		Problemtype:     problemtype,
		Status:          1,

		File_extension: ext,
		File_size:      int(size),
	}

	// database.Instance.Create(&problemrecord)
	// c.Set("Content-Type", "application/json")
	// return c.JSON(problemrecord)

	c.Set("Content-Type", "application/json")
	database.Instance.Create(&problemrecord)
	return c.JSON(problemrecord)

}


func GetProblemRecords(c *fiber.Ctx) error {
	var problemrecords []entities.ProblemRecord
	database.Instance.Find(&problemrecords)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecords)
	return nil
}



func GetProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var problemrecord entities.ProblemRecord
	database.Instance.First(&problemrecord, id)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)

	return nil
}


func UpdateProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	operator := c.FormValue("operator")
	var data_problem entities.ProblemRecord
	database.Instance.Where("id = ?",id).Find(&data_problem)
	fmt.Println(data_problem.ID)
	if id != data_problem.ID {
	
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}
	problemrecord := entities.ProblemRecord{
		Operator:        operator,
		Status:          2,

	}
	
	if err := c.BodyParser(&problemrecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if database.Instance.Where("id = ?", id).Updates(&problemrecord).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Error Update File",
		})
		}

	// database.Instance.Save(&problemrecord)
	return c.JSON(problemrecord)

	
}

func CompletedProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	Casuseproblem := c.FormValue("casuseproblem")
	Solution := c.FormValue("solution")
	Suggestion := c.FormValue("suggestion")
	var data_problem entities.ProblemRecord
	database.Instance.Where("id = ?",id).Find(&data_problem)
	fmt.Println(data_problem.ID)
	if id != data_problem.ID {
	
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}
	problemrecord := entities.ProblemRecord{
		Casuseproblem:        Casuseproblem,
		Solution:          Solution,
		Suggestion:          Suggestion,
		Status:          3,

	}
	
	if err := c.BodyParser(&problemrecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if database.Instance.Where("id = ?", id).Updates(&problemrecord).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Error Update File",
		})
		}


	//database.Instance.Where("id = ?",id).Save(&problemrecord)
	return c.JSON(problemrecord)
}


func DeleteProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var problemrecord entities.ProblemRecord
	database.Instance.First(&problemrecord, id)
	if problemrecord.ID == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}
	database.Instance.Delete(&problemrecord)
	return c.JSON(fiber.Map{
		"message": "Record deleted",
	})
}

func GetProblemRecordByAgency(c *fiber.Ctx) error {
	agency := c.Params("agency")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("agency = ?", agency).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByContact(c *fiber.Ctx) error {
	contact := c.Params("contact")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("contact = ?", contact).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblem(c *fiber.Ctx) error {
	problem := c.Params("problem")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problem = ?", problem).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByLevel(c *fiber.Ctx) error {
	level := c.Params("level")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("level = ?", level).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByInformer(c *fiber.Ctx) error {
	informer := c.Params("informer")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("informer = ?", informer).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByInformermessage(c *fiber.Ctx) error {
	informermessage := c.Params("informermessage")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("informermessage = ?", informermessage).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordBySystem(c *fiber.Ctx) error {
	system := c.Params("system")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("system = ?", system).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemtype(c *fiber.Ctx) error {
	problemtype := c.Params("problemtype")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemtype = ?", problemtype).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemstatus(c *fiber.Ctx) error {
	problemstatus := c.Params("problemstatus")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemstatus = ?", problemstatus).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemtime(c *fiber.Ctx) error {
	problemtime := c.Params("problemtime")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemtime = ?", problemtime).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemdescription(c *fiber.Ctx) error {
	problemdescription := c.Params("problemdescription")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemdescription = ?", problemdescription).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemimage(c *fiber.Ctx) error {
	problemimage := c.Params("problemimage")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemimage = ?", problemimage).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemvideo(c *fiber.Ctx) error {
	problemvideo := c.Params("problemvideo")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemvideo = ?", problemvideo).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemvoice(c *fiber.Ctx) error {
	problemvoice := c.Params("problemvoice")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemvoice = ?", problemvoice).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemfile(c *fiber.Ctx) error {
	problemfile := c.Params("problemfile")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemfile = ?", problemfile).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemlocation(c *fiber.Ctx) error {
	problemlocation := c.Params("problemlocation")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemlocation = ?", problemlocation).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemtimeend(c *fiber.Ctx) error {
	problemtimeend := c.Params("problemtimeend")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemtimeend = ?", problemtimeend).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}


// func (book *entities.ProblemRecord) BeforeCreate(tx *gorm.DB) (err error) {
//     book.ID = uuid.New().String()
// 	return
// }

// func CreateProblemRecord(w http.ResponseWriter, r *http.Request) {
// 	//var problemrecord entities.ProblemRecord
// 	//_ = json.NewDecoder(r.Body).Decode(&problemrecord)
// 	//database.Instance.Create(&problemrecord)
// 	//json.NewEncoder(w).Encode(problemrecord)
// 	r.ParseMultipartForm(10 * 1024 * 1024)
// 	file, handler, err := r.FormFile("problem_records")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	//get file name
// 	file_name := handler.Filename
// 	//get file extension
// 	ext := filepath.Ext(handler.Filename)
// 	//get file size
// 	size := handler.Size
// 	//get file path
// 	file_path := "/upload/" + file_name + ext
// 	fmt.Println("File Info")
// 	fmt.Println("File Name:", file_name)
// 	fmt.Println("File Size:", size)
// 	fmt.Println("File Path:", file_path)
// 	fmt.Println("File Extension:", ext)
// 	filename := path.Join("uploads", path.Base(handler.Filename))
// 	dest, err := os.Create(filename)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	if _, err = io.Copy(dest, file); err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer dest.Close()

// 	agency := r.FormValue("agency")
// 	contact := r.FormValue("contact")
// 	problem := r.FormValue("problem")
// 	level := r.FormValue("level")
// 	informer := r.FormValue("informer")
// 	informermessage := r.FormValue("informermessage")
// 	system := r.FormValue("system")
// 	problemtype := r.FormValue("problemtype")

// 	w.Header().Set("Content-Type", "application/json")

// 	problemrecord := entities.ProblemRecord{
// 		File_name:      file_name,
// 		Path_file:      file_path,
// 		Agency:         agency,
// 		Contact:        contact,
// 		Problem:        problem,
// 		Level:          level,
// 		Informer:       informer,
// 		Informermessage: informermessage,
// 		System:         system,
// 		Problemtype:   problemtype,

// 		File_extension: ext,
// 		File_size:      int(size),
// 	}

// 	database.Instance.Create(&problemrecord)
// 	json.NewEncoder(w).Encode(&problemrecord)
// }

// func GetProblemRecords(w http.ResponseWriter, r *http.Request) {
// 	var problemrecords []entities.ProblemRecord
// 	database.Instance.Find(&problemrecords)
// 	json.NewEncoder(w).Encode(problemrecords)
// 	var problemsender []entities.ProblemSender
// 	database.Instance.Find(&problemsender)
// 	json.NewEncoder(w).Encode(problemsender)
// }

// func GetProblemRecord(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var problemrecord entities.ProblemRecord
// 	database.Instance.First(&problemrecord, params["id"])
// 	json.NewEncoder(w).Encode(problemrecord)
// }
