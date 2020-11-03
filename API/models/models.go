package models

type Company struct {
	ID          int    `json:"Id"`
	Name        string `json:"Name"`
	Headquarter string `json:"Headquater"`
}

type Employee struct {
	ID        int    `json:"Id"`
	FName     string `json:"FName"`
	LName     string `json:"LName"`
	Profile   string `json:"Profile"`
	CompanyID int    `json:"CompanyId"`
}
