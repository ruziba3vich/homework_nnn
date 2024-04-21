package services

import (
	"database/sql"
	"fmt"

	"github.com/ruziba3vich/blog-app/internal/models"
)

func ApplyService(user models.User, vacancy models.Vacancy, db *sql.DB) error {
	err := SendEmailService(&user, &vacancy.Employer, GetApplicationContent(user, vacancy), db)
	if err != nil {
		return err
	}
	return nil
}
func GetApplicationContent(user models.User, vacancy models.Vacancy) string {
	letter := fmt.Sprintf(`Subject: Application for %s Position at %s

Dear Hiring Manager,
I hope this email finds you well. My name is %s, and I am writing to express my interest in the %s position at %s.

I am confident in my ability to contribute effectively to your team and help %s achieve its goals. Throughout my career, I have developed strong skills and demonstrated a commitment to excellence in everything I do.

I am particularly drawn to %s's, and I am eager to be part of a dynamic and forward-thinking organization like yours.

I have attached my resume for your review. I would welcome the opportunity to further discuss how my skills and experiences align with the requirements of the %s position. Please feel free to contact me at %s to schedule a meeting.

Thank you for considering my application. I look forward to the possibility of contributing to the success of %s.

Best regards,
%s`,
		vacancy.Title,
		vacancy.CompanyName,
		user.Fullname,
		vacancy.Title,
		vacancy.CompanyName,
		vacancy.CompanyName,
		vacancy.CompanyName,
		vacancy.Title,
		user.Email,
		vacancy.CompanyName,
		user.Fullname)

	return letter
}
