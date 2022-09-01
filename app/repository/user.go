package repository

import (
	"backend/pkg/models"
	"mime/multipart"
)

type UserRepository interface {
	Create(user *models.User) (uint, error)
	Update(user *models.User) error
	SaveAvatar(user *models.User, header *multipart.FileHeader) error
	IsAbleToLogin(email string, password string) (bool, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(IdU uint) (*models.User, error)
	IsExist(email string) (bool, error)
}
