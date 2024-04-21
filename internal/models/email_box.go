package models

type EmailBox struct {
	UserId     int    `json:"user_id"`
	EmployerId int    `json:"employer_id"`
	Contenet   string `json:"content"`
}
