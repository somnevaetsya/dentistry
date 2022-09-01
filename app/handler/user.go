package handler

import (
	"backend/app/usecase"
	customErrors "backend/pkg/errors"
	"backend/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"time"
)

type UserHandler struct {
	usecase    usecase.UserUseCase
	expireDays int64
}

func MakeUserHandler(usecase usecase.UserUseCase, exp int64) *UserHandler {
	return &UserHandler{usecase: usecase, expireDays: exp}
}

func (userHandler *UserHandler) Login(c *gin.Context) {
	var user models.User
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := jsoniter.Unmarshal(body, &user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	// вызываю юзкейс

	userId, token, err := userHandler.usecase.Login(user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	user, err = userHandler.usecase.GetInfoById(userId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	expiration := time.Now().Add(time.Duration(userHandler.expireDays) * 24 * time.Hour)
	fmt.Println("EXP:", expiration)
	cookie := http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiration,
		//HttpOnly: true,
		//SameSite: http.SameSiteLaxMode,
		Path: "/",
		//Secure:   true,
	}

	userJson, err := jsoniter.Marshal(user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	http.SetCookie(c.Writer, &cookie)
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (userHandler *UserHandler) Register(c *gin.Context) {
	var user models.User
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := jsoniter.Unmarshal(body, &user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	userId, token, err := userHandler.usecase.Register(user)

	if err != nil {
		_ = c.Error(err)
		return
	}

	user, err = userHandler.usecase.GetInfoById(userId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	expiration := time.Now().Add(time.Duration(userHandler.expireDays) * 24 * time.Hour)
	cookie := http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiration,
		//HttpOnly: true,
		//SameSite: http.SameSiteLaxMode,
		Path: "/",
		//Secure:   true,
	}

	userJson, err := jsoniter.Marshal(user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	http.SetCookie(c.Writer, &cookie)
	c.Data(http.StatusCreated, "application/json; charset=utf-8", userJson)
}

func (userHandler *UserHandler) Logout(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	err = userHandler.usecase.Logout(token)

	if err != nil {
		_ = c.Error(err)
		return
	}

	var resp models.Response
	resp.Desc = "Succecfully logout"
	respJson, err := jsoniter.Marshal(resp)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.SetCookie("token", token, -1, "", "", false, true)
	c.Data(http.StatusOK, "application/json; charset=utf-8", respJson)
}

func (userHandler *UserHandler) GetInfoByCookie(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	user, err := userHandler.usecase.GetInfoById(uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	userJson, err := jsoniter.Marshal(user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (userHandler *UserHandler) SaveAvatar(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	header, err := c.FormFile("avatar")
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	user := new(models.User)
	user.IdU = uint(userId.(uint64))
	user.ImgAvatar = header.Filename

	path, err := userHandler.usecase.SaveAvatar(user, header)

	if err != nil {
		_ = c.Error(err)
		return
	}

	var avatarPath models.Avatar
	avatarPath.AvatarPath = path
	avatarPathJson, err := jsoniter.Marshal(avatarPath)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", avatarPathJson)

}

func (userHandler *UserHandler) RefactorProfile(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var user models.User
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := jsoniter.Unmarshal(body, &user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	user.IdU = uint(userId.(uint64))
	err = userHandler.usecase.RefactorProfile(user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var resp models.Response
	resp.Desc = "Profile successfully updated"
	respJson, err := jsoniter.Marshal(resp)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", respJson)
}

func (userHandler *UserHandler) VerificationProfile(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var code models.Code
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := jsoniter.Unmarshal(body, &code)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	code.UserId = userId.(uint64)
	err = userHandler.usecase.CheckCode(code)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var resp models.Response
	resp.Desc = "Profile successfully verificated"
	respJson, err := jsoniter.Marshal(resp)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", respJson)
}
