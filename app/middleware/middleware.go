package middleware

import (
	"backend/app/repository"
	customErrors "backend/pkg/errors"
	"backend/pkg/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type Middleware struct {
	sessionRepo repository.SessionRepository
}

func CreateMiddleware(sessionRepo repository.SessionRepository) *Middleware {
	return &Middleware{sessionRepo: sessionRepo}
}

func (mw *Middleware) CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	// Получаю сессии из БД
	userId, err := mw.sessionRepo.GetSession(token)
	if err != nil {
		return
	}
	c.Set("Auth", userId)
}

func CheckError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := errors.Unwrap(c.Errors.Last())
			newErr := new(models.CustomError)
			newErr.CustomErr = err.Error()
			errJson, err1 := jsoniter.Marshal(newErr)
			if err1 != nil {
				fmt.Println(err1.Error())
				return
			}
			c.Data(customErrors.ConvertErrorToCode(err), "application/json; charset=utf-8", errJson)
			return
		}
	}
}
