package mail

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, htmlBody string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "MediLink Security <"+os.Getenv("SMTP_EMAIL")+">")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT")) // Default 587

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"), // smtp.gmail.com
		port,
		os.Getenv("SMTP_EMAIL"),    // email-anda@gmail.com
		os.Getenv("SMTP_PASSWORD"), // APP PASSWORD 16 digit (bukan password login!)
	)

	return d.DialAndSend(m)
}
