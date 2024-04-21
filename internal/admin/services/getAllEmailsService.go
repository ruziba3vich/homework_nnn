package services

import "github.com/ruziba3vich/blog-app/internal/models"

func GetAllEmailsService(employer models.User) []models.EmailBox {
	return employer.EmailBox
}
