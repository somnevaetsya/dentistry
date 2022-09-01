package impl

import (
	"backend/app/repository"
	"backend/cmd/email/handler"
	"backend/pkg/models"
	"context"
)

type EmailRepositoryImpl struct {
	client handler.EmailCheckerClient
	ctx    context.Context
}

func CreateEmailRepo(cl handler.EmailCheckerClient) repository.EmailRepository {
	return &EmailRepositoryImpl{client: cl, ctx: context.Background()}
}

func (redisConnect *EmailRepositoryImpl) SendEmail(email models.Email) error {
	_, err := redisConnect.client.SendEmail(redisConnect.ctx, &handler.EmailModel{Email: email.Address, Code: email.Code})
	return err
}
