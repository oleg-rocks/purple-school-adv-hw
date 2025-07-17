package sender

import (
	"fmt"
	"go/adv-hw/configs"
	"net/smtp"
	"github.com/jordan-wright/email"
)

func SendVerification(toEmail string, hash string, config *configs.Config) error {
	e := email.NewEmail()
	e.From = "mean.elvis77@gmail.com"
	e.To = []string{toEmail}
	e.Subject = "Email Verification"
	message := fmt.Sprintf("Please verify your email by clicking the following link: http://localhost:8081/verify/%s", hash)
	e.Text = []byte(message)
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	return e.Send("sandbox.smtp.mailtrap.io:587", auth)
}