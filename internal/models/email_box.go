package models

import (
	"time"
)

type EmailBox struct {
	Id      int       `json:"id"`
	To      User      `json:"user_id"`
	From    User      `json:"employer_id"`
	Content string    `json:"content"`
	SentOn  time.Time `json:"sent_on"`
}
