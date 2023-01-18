//not use
package problemrecord

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
)

func Uploadfile(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(10 * 1024 * 1024)
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
	path_file := "./upload/" + file_name 
	fmt.Println("File Info")
	// fmt.Println("File Name:", handler.Filename)
	// fmt.Println("File Size:", handler.Size)
	// fmt.Println("File Type:", handler.Header.Get("Content-Type"))
	fmt.Println("File Name:", file_name)
	fmt.Println("File Size:", size)
	fmt.Println("File Path:", path_file)
	fmt.Println("File Extension:", ext)
	filename := path.Join("uploads", path.Base(handler.Filename))
	dest, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Servers Error", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	database.Instance.Create(&entities.ProblemRecord{
		File_name:      file_name,
		Path_file:      path_file,
		File_extension: ext,
		File_size:      int(size),
	})
	json.NewEncoder(w).Encode("Upload File Successfully!")
}