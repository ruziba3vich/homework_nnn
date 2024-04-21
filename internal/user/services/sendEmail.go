package services

import (
	"database/sql"
	"time"

	"github.com/ruziba3vich/blog-app/internal/models"
)

func SendEmailService(sender *models.User, reciepent *models.User, content string, db *sql.DB) error {
	newEmail := models.EmailBox{
		From:    *sender,
		To:      *reciepent,
		Content: content,
		SentOn:  time.Now(),
	}
	query := "INSERT INTO Email_box(user_id, employer_id, content, sent_on) RETURNING id;"
	err := db.QueryRow(query, newEmail.To.ID, newEmail.From.ID, content, newEmail.SentOn).Scan(&newEmail.Id)
	if err != nil {
		return err
	}
	reciepent.EmailBox = append(reciepent.EmailBox, newEmail)

	return nil
}
