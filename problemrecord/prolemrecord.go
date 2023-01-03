package problemrecord

import (
	"encoding/json"
	"fmt"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"
)

//use form-data

func CreateProblemRecord(w http.ResponseWriter, r *http.Request) {

	//var problemrecord entities.ProblemRecord
	//_ = json.NewDecoder(r.Body).Decode(&problemrecord)
	//database.Instance.Create(&problemrecord)
	//json.NewEncoder(w).Encode(problemrecord)
	r.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := r.FormFile("problem_records")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//get file name
	file_name := handler.Filename
	//get file extension
	ext := filepath.Ext(handler.Filename)
	//get file size
	size := handler.Size
	//get file path
	file_path := "/upload/" + file_name + ext
	fmt.Println("File Info")
	fmt.Println("File Name:", file_name)
	fmt.Println("File Size:", size)
	fmt.Println("File Path:", file_path)
	fmt.Println("File Extension:", ext)
	filename := path.Join("uploads", path.Base(handler.Filename))
	dest, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dest.Close()
	agency := r.FormValue("agency")
	contact := r.FormValue("contact")
	problem := r.FormValue("problem")
	level := r.FormValue("level")
	informer := r.FormValue("informer")
	informermessage := r.FormValue("informermessage")
	system := r.FormValue("system")
	problemtype := r.FormValue("problemtype")
	w.Header().Set("Content-Type", "application/json")

	problemrecord := entities.ProblemRecord{
		File_name:      file_name,
		Path_file:      file_path,
		Agency:         agency,
		Contact:        contact,
		Problem:        problem,
		Level:          level,
		Informer:       informer,
		Informermessage: informermessage,
		System:         system,
		Problemtype:   problemtype,

		File_extension: ext,
		File_size:      int(size),
	}

	database.Instance.Create(&problemrecord)
	json.NewEncoder(w).Encode(&problemrecord)
}

func GetProblemRecords(w http.ResponseWriter, r *http.Request) {
	var problemrecords []entities.ProblemRecord
	database.Instance.Find(&problemrecords)
	json.NewEncoder(w).Encode(problemrecords)
}

func GetProblemRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var problemrecord entities.ProblemRecord
	database.Instance.First(&problemrecord, params["id"])
	json.NewEncoder(w).Encode(problemrecord)
}
