package services

import (
	"database/sql"

	"github.com/ruziba3vich/blog-app/internal/authentication"
	"github.com/ruziba3vich/blog-app/internal/models"
)

func Register(user models.User, db *sql.DB) (string, error) {
	query := "INSERT INTO users(username, fullname, email, country, city) VALUES ($1, $2, $3, $4, $5);"
	var erro error
	user.Password, erro = authentication.HashPassword(user.Password)
	if erro != nil {
		return "", erro
	}
	if _, err := db.Exec(query, user.Username, user.Fullname, user.Email, user.Country, user.City); err != nil {
		return "", err
	}
	db.QueryRow("SELECT id FROM Users WHERE username = $1", user.Username).Scan(&user.ID)
	token, err := authentication.GenerateJWTToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
