package entities

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Nickname string `json:"nickname"`
	Systems string `json:"systems"`
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
	ID   int    `json:"id"`
	Agency string `json:"agency"`
	Contact string `json:"contact"`
	Informer string `json:"informer"`
	Informermessage string `json:"informermessage"`
	System string `json:"system"`
	Problemtype string `json:"problemttype"`
	Level string `json:"level"`
	Problem string `json:"problem"`
	file string `json:"file"`
	
}

