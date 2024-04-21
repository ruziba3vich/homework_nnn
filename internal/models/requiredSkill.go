package models

type RequiredSkills struct {
	Id        int `json:"id"`
	VacancyId int `json:"vacancy_id"`
	SkillId   int `json:"skill_id"`
}
