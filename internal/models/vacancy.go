package models

type Vacancy struct {
	Id          int    `json:"id"`
	CompanyName string `json:"company_name"`
	Employer    User   `json:"employer_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Country     string `json:"country"`
	City        string `json:"city"`
	IsActive    bool   `json:"is_active"`
	Position    string `json:"position"`
}
