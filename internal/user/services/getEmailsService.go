package services

import "github.com/ruziba3vich/blog-app/internal/models"

func GetAllEmailsService(user models.User) []models.EmailBox {
	return user.EmailBox
}
