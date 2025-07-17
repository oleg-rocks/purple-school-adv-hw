package sender

import (
	"fmt"
	"go/adv-hw/configs"
	"net/smtp"
	"github.com/jordan-wright/email"
)

func SendVerification(toEmail string, hash string, config *configs.Config) error {
	e := email.NewEmail()
	e.From = config.Email
	e.To = []string{toEmail}
	e.Subject = "Email Verification"
	message := fmt.Sprintf("Please verify your email by clicking the following link: http://localhost:8081/verify/%s", hash)
	e.Text = []byte(message)
	auth := smtp.PlainAuth("", "cb2bca2dc665f2", config.Password, config.Address)
	return e.Send("sandbox.smtp.mailtrap.io:587", auth)
}