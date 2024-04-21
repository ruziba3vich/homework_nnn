package models

type Skill struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	From User
	To   User
}
