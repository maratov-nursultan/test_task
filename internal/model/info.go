package model

type ErrorInfo struct {
	Success bool   `json:"success"`
	Errors  string `json:"errors"`
}

func NewInfo(success bool, errors string) *ErrorInfo {
	return &ErrorInfo{
		Success: success,
		Errors:  errors,
	}
}
