package models

type Employer struct {
	ID          int    `json:"id"`
	Fullname    string `json:"fullname"`
	CompanyName string `json:"company_name"`
}
