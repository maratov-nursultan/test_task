package model

type IinCheckResponse struct {
	Correct     bool   `json:"correct"`
	Sex         string `json:"sex"`
	DateOfBirth string `json:"date_of_birth"`
}
