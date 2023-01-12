package entities

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
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
	ID              string    `gorm:"primaryKey" json:"id"` 
	Agency          string    `gorm:"type:varchar(50)" json:"agency"`
	Contact         string    `gorm:"type:varchar(50)" json:"contact"`
	Informer        string    `gorm:"type:varchar(50)" json:"informer"`
	Informermessage string    `gorm:"type:varchar(50)" json:"informermessage"`
	System          string    `gorm:"type:varchar(50)" json:"system"`
	//System 		     *System  `gorm:"foreignKey:SystemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"system"`
	Problemtype     string    `gorm:"type:varchar(50)" json:"problemttype"`
	Level           string    `gorm:"type:varchar(50)" json:"level"`
	Problem         string    `gorm:"type:varchar(50)" json:"problem"`
	File_name       string `json:"file_name"`
	Path_file       string `json:"path_file"`
	File_extension  string `json:"file_extension"`
	File_size       int    `json:"file_size"`
	Status          string   `json:"status"`
	Casuseproblem string    `gorm:"type:varchar(255)" json:"casuseproblem"`
	Solution      string   	`gorm:"type:varchar(255)" json:"solution"`
	Suggestion    string    `gorm:"type:varchar(255)" json:"suggestion"`
	Operator        string  `gorm:"type:varchar(50)" json:"operator"`

	CreatedAt       time.Time `gorm:"<-:crate;type:timestamp" json:"created_at"`
	//SenderAt     	time.Time `gorm:"<-:update;type:timestamp" json:"sender_at"`
	SenderAt time.Time `gorm:"column:sender_at;type:TIMESTAMP;DEFAULT:CURRENT_TIMESTAMP;not null;" json:"sender_at"`
	// UpdatedAt time.Time `gorm:"<-:update;type:timestamp;column:sender_at" json:"updated_at"`
	CompletedAt     time.Time `gorm:"column:completed_at;type:TIMESTAMP;DEFAULT:CURRENT_TIMESTAMP;not null;" json:"completed_at"`

}

func (book *ProblemRecord) BeforeCreate(tx *gorm.DB) (err error) {
	uuidWithHyphen := uuid.New()
	//limit string length to 32 characters
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	book.ID = uuid
	return
}

type TimeWork struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Operator      string    `gorm:"type:varchar(255)" json:"operator"`
	Worktime      string    `gorm:"type:varchar(255)" json:"worktime"`
}


 type Status struct {
	ID  int    `json:"id"`
	Name string `json:"name"`
}





