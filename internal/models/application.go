package models

type Application struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	VacancyId int `json:"vacancy_id"`
}
