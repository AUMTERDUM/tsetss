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


	if err := c.BodyParser(&problemrecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := database.Instance.Create(&problemrecord).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(problemrecord)
}

func GetProblemRecords(c *fiber.Ctx) {
	var problemrecords []entities.ProblemRecord
	database.Instance.Find(&problemrecords)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecords)

	var problemsender []entities.ProblemSender
	database.Instance.Find(&problemsender)
	c.JSON(problemsender)
}

func GetProblemRecord(c *fiber.Ctx) {
	id := c.Params("id")
	var problemrecord entities.ProblemRecord
	database.Instance.First(&problemrecord, id)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
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
