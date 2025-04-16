package model

type InfoRequest struct {
	Name  string `json:"name"`
	Iin   string `json:"iin"`
	Phone string `json:"phone"`
}

type ListUserByNameRequest struct {
	Firstname  string
	Lastname   string
	Middlename string
}

type User struct {
	Name  string `json:"name"`
	Iin   string `json:"iin"`
	Phone string `json:"phone"`
}
