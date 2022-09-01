package impl

import (
	"backend/app/repository"
	"backend/app/usecase"
	customErrors "backend/pkg/errors"
	"backend/pkg/hash"
	"backend/pkg/models"
	"backend/pkg/token"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
)

type UserUseCaseImpl struct {
	rep   repository.UserRepository
	red   repository.SessionRepository
	email repository.EmailRepository
}

func MakeUserUsecase(rep_ repository.UserRepository, red_ repository.SessionRepository, email_ repository.EmailRepository) usecase.UserUseCase {
	return &UserUseCaseImpl{rep: rep_, red: red_, email: email_}
}

func (userUseCase *UserUseCaseImpl) Login(user models.User) (uint, string, error) {
	// вызываю из бд проверку есть ли юзер
	//сравниваю пароли
	isAble, err := userUseCase.rep.IsAbleToLogin(user.Email, user.Password)
	if err != nil {
		return 0, "", err
	}
	if !isAble {
		return 0, "", customErrors.ErrUnauthorized
	}

	newUser, err := userUseCase.rep.GetUserByEmail(user.Email)
	if err != nil {
		return 0, "", err
	}

	newToken := token.GenerateSessionToken()
	err = userUseCase.red.SetSession(models.Session{UserId: newUser.IdU, CookieValue: newToken})
	// сохраняю сессию в бд и возвращаю token
	if err != nil {
		return 0, "", err
	}
	return newUser.IdU, newToken, nil
}

func (userUseCase *UserUseCaseImpl) Register(user models.User) (uint, string, error) {
	err := hash.CheckPassword(user.Password)
	if err != nil {
		return 0, "", err
	}

	isValidEmail := hash.CheckEmail(user.Email)
	if !isValidEmail {
		return 0, "", customErrors.ErrEmailNotValid
	}

	// проверяю в БД есть ли такой юзер и обрабатываю ошибку в случае чего

	isExist, err := userUseCase.rep.IsExist(user.Email)
	if isExist {
		return 0, "", customErrors.ErrEmailExist
	} else if err != nil {
		return 0, "", err
	}

	hashPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return 0, "", err
	}
	//задаем текущему пользователю "новый" пароль
	user.Password = hashPassword

	// добавляю юзера в бд и создаю токен для него, добавляю в бд сессию

	userId, err := userUseCase.rep.Create(&user)
	if err != nil {
		return 0, "", err
	}

	currToken := token.GenerateSessionToken()
	err = userUseCase.red.SetSession(models.Session{UserId: userId, CookieValue: currToken})
	if err != nil {
		return 0, "", err
	}

	code := hash.EncodeToString(6)
	err = userUseCase.email.SendEmail(models.Email{Code: code, Address: user.Email})
	if err != nil {
		return 0, "", err
	}

	err = userUseCase.red.SetCode(models.Code{UserId: uint64(userId), Code: code})
	// возвращаю токен и ошибку
	return userId, currToken, err
}

func (userUseCase *UserUseCaseImpl) Logout(token string) error {
	err := userUseCase.red.DeleteSession(token)
	return err
}

func (userUseCase *UserUseCaseImpl) GetInfoById(userId uint) (models.User, error) {
	// получаю из бд всю инфу по айдишнику кроме пароля
	user, err := userUseCase.rep.GetUserById(userId)
	if err != nil {
		return models.User{}, err
	}

	user.Password = ""
	return *user, err
}

func (userUseCase *UserUseCaseImpl) SaveAvatar(user *models.User, header *multipart.FileHeader) (string, error) {
	err := userUseCase.rep.SaveAvatar(user, header)
	return strings.Join([]string{strconv.Itoa(int(user.IdU)), ".webp"}, ""), err
}

func (userUseCase *UserUseCaseImpl) RefactorProfile(user models.User) error {
	err := hash.CheckPassword(user.Password)
	if err != nil {
		return err
	}

	isValidEmail := hash.CheckEmail(user.Email)
	if !isValidEmail {
		return customErrors.ErrEmailNotValid
	}
	return userUseCase.rep.Update(&user)
}

func (userUseCase *UserUseCaseImpl) CheckCode(code models.Code) error {
	trueCode, err := userUseCase.red.GetCode(code.UserId)
	fmt.Println(trueCode)
	if err != nil || trueCode != code.Code {
		return customErrors.ErrInvalidCode
	} else {
		currUser, err := userUseCase.rep.GetUserById(uint(code.UserId))
		if err != nil {
			return err
		}
		currUser.IsConfirmed = true
		err = userUseCase.rep.Update(currUser)
		if err != nil {
			return err
		}
		return nil
	}
}
