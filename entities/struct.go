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
