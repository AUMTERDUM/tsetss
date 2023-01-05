package entities

import (
	"time"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Nickname string `gorm:"type:varchar(50)" json:"nickname"`
	Systems  string `gorm:"type:varchar(50)" json:"systems"`
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
	ID              int    `gorm:"primaryKey" json:"id"`
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
	CreatedAt       time.Time `gorm:"<-:create;type:timestamp;" json:"created_at"`

}

type ProblemSender struct {
	ID              int    	`gorm:"primaryKey" json:"id"`
	Operator        string    `json:"operator"`
	Sender_At     time.Time `gorm:"<-:update;type:timestamp;" json:"Sender_at"`
}

type CompleteRecord struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Casuseproblem string    `gorm:"type:varchar(255)" json:"casuseproblem"`
	Solution      string   	`gorm:"type:varchar(255) "json:"solution"`
	Suggestion    string    `gorm:"type:varchar(255)" json:"suggestion"`
	Completed_at     time.Time `gorm:"<-:update;type:timestamp;" json:"completed_at"`
}

// type File struct {
// 	Id             int    `json:"id"`
// 	File_name      string `json:"file_name"`
// 	Path_file      string `json:"path_file"`
// 	File_extension string `json:"file_extension"`
// 	File_size      int    `json:"file_size"`
// 	Status         int    `json:"status"`
// }

