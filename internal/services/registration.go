package services

import (
	"database/sql"

	"github.com/ruziba3vich/blog-app/internal/authentication"
	"github.com/ruziba3vich/blog-app/internal/models"
)

func Register(user models.User, db *sql.DB) (string, error) {
	hashedPassword, err := authentication.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	query := "INSERT INTO users(username, fullname, email, password, country, city) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var userID int
	err = db.QueryRow(query, user.Username, user.Fullname, user.Email, hashedPassword, user.Country, user.City).Scan(&userID)
	if err != nil {
		return "", err
	}

	token, err := authentication.GenerateJWTToken(userID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
