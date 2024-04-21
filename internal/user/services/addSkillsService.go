package services

import (
	"database/sql"
	"sync"

	"github.com/ruziba3vich/blog-app/internal/models"
)

func AddSkillsService(user models.User, skills []string, db *sql.DB) error {
	var resultingError error
	var wg sync.WaitGroup
	for _, skill := range skills {
		if models.Skills[skill] == 0 {
			wg.Add(1)
			go func(skill string) {
				defer wg.Done()
				newSkill, err := AddSkillIntoDatabase(skill, db)
				if err != nil {
					resultingError = err
					return
				}
				models.AddNewValue(skill, newSkill.Id)
			}(skill)
		}
	}
	wg.Wait()

	return resultingError
}

func AddSkillIntoDatabase(skill string, db *sql.DB) (newSkill models.Skill, e error) {
	query := "INSERT INTO Skills(name) VALUES($1) RETURNING id"
	if err := db.QueryRow(query, skill).Scan(&newSkill.Id); err != nil {
		return newSkill, err
	}
	newSkill.Name = skill
	return newSkill, nil
}
