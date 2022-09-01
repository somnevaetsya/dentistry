package usecase

import (
	"backend/cmd/email/repository"
	"backend/pkg/models"
	"gopkg.in/gomail.v2"
)

type EmailUseCase struct {
	emailRepo    repository.EmailRepository
	mailUsername string
}

func CreateEmailUseCase(emailRepository repository.EmailRepository, user string) EmailUseCase {
	return EmailUseCase{emailRepo: emailRepository, mailUsername: user}
}

func (emailUseCase *EmailUseCase) SendEmail(email models.Email) error {
	emailLetter := gomail.NewMessage()
	emailLetter.SetHeader("From", emailUseCase.mailUsername)
	emailLetter.SetHeader("To", email.Address)
	emailLetter.SetHeader("Subject", "Добро пожаловать в сайт стоматологии!")
	emailLetter.SetBody("text/plain", "Рады вас видеть у себя на сайте! Код верификации:"+email.Code)
	return emailUseCase.emailRepo.SendEmail(emailLetter)
}
