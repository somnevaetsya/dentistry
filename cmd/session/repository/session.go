package repository

import (
	"backend/pkg/models"
	"github.com/go-redis/redis"
	"strconv"
)

var (
	expireCode = 86_400
)

type SessionRepository struct {
	client *redis.Client
}

func CreateSessRepo(cl *redis.Client) SessionRepository {
	return SessionRepository{client: cl}
}

func (redisConnect *SessionRepository) SetSession(session models.Session, cookieTime int) error {
	return redisConnect.client.Do("SETEX", session.CookieValue, cookieTime, session.UserId).Err()
}

func (redisConnect *SessionRepository) GetSession(cookieValue string) (uint64, error) {
	value, err := redisConnect.client.Get(cookieValue).Uint64()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (redisConnect *SessionRepository) DeleteSession(cookieValue string) error {
	err := redisConnect.client.Del(cookieValue).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redisConnect *SessionRepository) SetCode(code models.Code) error {
	return redisConnect.client.Do("SETEX", strconv.Itoa(int(code.UserId)), expireCode, code.Code).Err()
}

func (redisConnect *SessionRepository) GetCode(userId uint64) (string, error) {
	value, err := redisConnect.client.Get(strconv.FormatUint(userId, 10)).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (redisConnect *SessionRepository) DeleteCode(userId uint64) error {
	err := redisConnect.client.Del(strconv.FormatUint(userId, 10)).Err()
	if err != nil {
		return err
	}
	return nil
}
