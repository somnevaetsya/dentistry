package impl

import (
	"backend/cmd/session/handler"
	"backend/cmd/session/usecase"
	customErrors "backend/pkg/errors"
	"backend/pkg/models"
	"context"
)

type SessionServerImpl struct {
	sessUseCase usecase.SessionUseCase
	handler.UnimplementedSessionCheckerServer
	expire int
}

func CreateSessionServer(sessUseCase usecase.SessionUseCase, cookieExpiration int) handler.SessionCheckerServer {
	return &SessionServerImpl{sessUseCase: sessUseCase, expire: cookieExpiration}
}

func (sessServer *SessionServerImpl) Create(ctx context.Context, in *handler.SessionModel) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}
	sess := models.Session{UserId: uint(in.Userid), CookieValue: in.CookieValue}

	err := sessServer.sessUseCase.SetSession(sess, sessServer.expire)
	if err != nil {
		return &handler.Nothing{}, err
	}
	return &handler.Nothing{}, nil
}

func (sessServer *SessionServerImpl) Get(ctx context.Context, in *handler.SessionValue) (*handler.UserId, error) {
	if in == nil {
		return &handler.UserId{Id: 0}, customErrors.ErrBadInputData
	}
	userId, err := sessServer.sessUseCase.GetSession(in.CookieValue)
	if err != nil {
		return &handler.UserId{Id: 0}, err
	}
	return &handler.UserId{Id: userId}, nil
}

func (sessServer *SessionServerImpl) Delete(ctx context.Context, in *handler.SessionValue) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}
	err := sessServer.sessUseCase.DeleteSession(in.CookieValue)
	if err != nil {
		return &handler.Nothing{}, err
	}
	return &handler.Nothing{}, nil
}

func (sessServer *SessionServerImpl) CreateCode(ctx context.Context, in *handler.CodeModel) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}
	code := models.Code{UserId: in.Userid, Code: in.Code}

	err := sessServer.sessUseCase.SetCode(code)
	if err != nil {
		return &handler.Nothing{}, err
	}
	return &handler.Nothing{}, nil
}

func (sessServer *SessionServerImpl) GetCode(ctx context.Context, in *handler.UserId) (*handler.Code, error) {
	if in == nil {
		return &handler.Code{Code: ""}, customErrors.ErrBadInputData
	}
	code, err := sessServer.sessUseCase.GetCode(in.Id)
	if err != nil {
		return &handler.Code{Code: ""}, err
	}
	return &handler.Code{Code: code}, nil
}

func (sessServer *SessionServerImpl) DeleteCode(ctx context.Context, in *handler.UserId) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}
	err := sessServer.sessUseCase.DeleteCode(in.Id)
	if err != nil {
		return &handler.Nothing{}, err
	}
	return &handler.Nothing{}, nil
}
