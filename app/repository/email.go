package repository

import "backend/pkg/models"

type EmailRepository interface {
	SendEmail(email models.Email) error
}
