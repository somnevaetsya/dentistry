package impl

import (
	"backend/cmd/email/handler"
	"backend/cmd/email/usecase"
	customErrors "backend/pkg/errors"
	"backend/pkg/models"
	"context"
)

type EmailServerImpl struct {
	emailUseCase usecase.EmailUseCase
	handler.UnimplementedEmailCheckerServer
	expire int
}

func CreateEmailServer(emailUseCase usecase.EmailUseCase) handler.EmailCheckerServer {
	return &EmailServerImpl{emailUseCase: emailUseCase}
}

func (emailServer *EmailServerImpl) SendEmail(ctx context.Context, in *handler.EmailModel) (*handler.NothingSec, error) {
	if in == nil {
		return &handler.NothingSec{}, customErrors.ErrBadInputData
	}
	email := models.Email{Code: in.Code, Address: in.Email}

	err := emailServer.emailUseCase.SendEmail(email)
	if err != nil {
		return &handler.NothingSec{}, err
	}
	return &handler.NothingSec{}, nil
}
