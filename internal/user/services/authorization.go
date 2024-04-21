package services

import (
	"database/sql"
	"errors"

	"github.com/ruziba3vich/blog-app/internal/authentication"
)

func Authorize(username, password, role string, db *sql.DB) (string, error) {
	hashedPassword, err := authentication.HashPassword(password)
	if err != nil {
		return "", err
	}

	row := db.QueryRow("SELECT id FROM Users WHERE username = $1 AND password = $2", username, hashedPassword)
	var id int
	if err := row.Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user does not exist or invalid credentials")
		}
		return "", err
	}

	token, err := authentication.GenerateJWTToken(id, username, role)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
