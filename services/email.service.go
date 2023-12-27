package services

import (
	"fmt"

	"github.com/AkifhanIlgaz/key-aero-api/cfg"
	"github.com/AkifhanIlgaz/key-aero-api/models"
	"github.com/go-mail/mail/v2"
)

const DefaultSender = "keyaero@gmail.com"

type EmailService struct {
	dialer *mail.Dialer
}

func NewEmailService(config *cfg.Config) *EmailService {
	return &EmailService{
		dialer: mail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUsername, config.SMTPPassword),
	}
}

func (service *EmailService) Send(email models.Email) error {
	msg := mail.NewMessage()

	msg.SetHeader("To", email.To)
	msg.SetHeader("From", email.From)
	msg.SetBody("text/html", email.HTML)

	err := service.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

// TODO: Update HTML
// Link to update password page or login ?
func (service *EmailService) NewUser(to, username, password string) error {
	email := models.Email{
		From:    DefaultSender,
		To:      to,
		Subject: "User created",
		HTML: fmt.Sprintf(`  <div>
      <p>Username: %v</p>
      <p>Password: %v</p>
    </div>`, username, password),
	}

	err := service.Send(email)
	if err != nil {
		return fmt.Errorf("forgot password email: %w", err)
	}
	return nil
}
