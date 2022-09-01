package repository

import "backend/pkg/models"

type SessionRepository interface {
	SetSession(session models.Session) error
	GetSession(cookieValue string) (uint64, error)
	DeleteSession(cookieValue string) error
	SetCode(code models.Code) error
	GetCode(userId uint64) (string, error)
	DeleteCode(userId uint64) error
}
