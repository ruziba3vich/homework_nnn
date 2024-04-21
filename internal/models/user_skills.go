package models

type UserSkills struct {
	Id      int `json:"id"`
	UserId  int `json:"user_id"`
	SkillId int `json:"skill_id"`
}
