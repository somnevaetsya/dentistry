package impl

import (
	"backend/app/repository"
	"backend/cmd/session/handler"
	"backend/pkg/models"
	"context"
)

type SessionRepositoryImpl struct {
	client handler.SessionCheckerClient
	ctx    context.Context
}

func CreateSessionRepo(cl handler.SessionCheckerClient) repository.SessionRepository {
	return &SessionRepositoryImpl{client: cl, ctx: context.Background()}
}

func (redisConnect *SessionRepositoryImpl) SetSession(session models.Session) error {
	_, err := redisConnect.client.Create(redisConnect.ctx, &handler.SessionModel{CookieValue: session.CookieValue, Userid: uint64(session.UserId)})
	return err
}

func (redisConnect *SessionRepositoryImpl) GetSession(cookieValue string) (uint64, error) {
	userId, err := redisConnect.client.Get(redisConnect.ctx, &handler.SessionValue{CookieValue: cookieValue})
	return userId.Id, err
}

func (redisConnect *SessionRepositoryImpl) DeleteSession(cookieValue string) error {
	_, err := redisConnect.client.Delete(redisConnect.ctx, &handler.SessionValue{CookieValue: cookieValue})
	return err
}

func (redisConnect *SessionRepositoryImpl) SetCode(code models.Code) error {
	_, err := redisConnect.client.CreateCode(redisConnect.ctx, &handler.CodeModel{Userid: uint64(code.UserId), Code: code.Code})
	return err
}

func (redisConnect *SessionRepositoryImpl) GetCode(userId uint64) (string, error) {
	code, err := redisConnect.client.GetCode(redisConnect.ctx, &handler.UserId{Id: userId})
	return code.Code, err
}

func (redisConnect *SessionRepositoryImpl) DeleteCode(userId uint64) error {
	_, err := redisConnect.client.DeleteCode(redisConnect.ctx, &handler.UserId{Id: userId})
	return err
}
