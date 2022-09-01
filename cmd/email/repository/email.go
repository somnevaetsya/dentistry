package repository

import (
	"gopkg.in/gomail.v2"
)

var (
	expireCode = 86_400
)

type EmailRepository struct {
	dialer *gomail.Dialer
}

func CreateEmailRepo(d *gomail.Dialer) EmailRepository {
	return EmailRepository{dialer: d}
}

func (emailRepo *EmailRepository) SendEmail(message *gomail.Message) error {
	return emailRepo.dialer.DialAndSend(message)
}
