package entities

import (
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Systems  string `json:"systems"`
}

type System struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Problem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Level struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Time int    `json:"time"`
}

type Contact struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Agency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProblemRecord struct {
	ID              int       `json:"id"`
	Agency          string    `json:"agency"`
	Contact         string    `json:"contact"`
	Informer        string    `json:"informer"`
	Informermessage string    `json:"informermessage"`
	System          string    `json:"system"`
	Problemtype     string    `json:"problemttype"`
	Level           string    `json:"level"`
	Problem         string    `json:"problem"`
	File_name       string `json:"file_name"`
	Path_file       string `json:"path_file"`
	File_extension  string `json:"file_extension"`
	File_size       int    `json:"file_size"`
	Status          int    `json:"status"`

}

type ProblemSender struct {
	ID         int       `json:"id"`
	process    string    `json:"process"`
	process_at time.Time `json:"process_at"`
}

type CompleteRecord struct {
	ID            int       `json:"id"`
	casuseproblem string    `json:"casuseproblem"`
	solution      string    `json:"solution"`
	suggestion    string    `json:"suggestion"`
	complete_at   time.Time `json:"complete_at"`
}

// type File struct {
// 	Id             int    `json:"id"`
// 	File_name      string `json:"file_name"`
// 	Path_file      string `json:"path_file"`
// 	File_extension string `json:"file_extension"`
// 	File_size      int    `json:"file_size"`
// 	Status         int    `json:"status"`
// }