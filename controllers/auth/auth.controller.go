package controllers

import (
	"net/http"
	"strconv"

	"github.com/IqbalLx/alterra-agmc/entities"
	"github.com/IqbalLx/alterra-agmc/errors"
	services "github.com/IqbalLx/alterra-agmc/services/auth"
	"github.com/IqbalLx/alterra-agmc/utils"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService services.IAuthService
}

func NewAuthController(userService services.IAuthService) *AuthController {
	return &AuthController{userService}
}

func (ac AuthController) Register(c echo.Context) error {
	newUser := entities.User{}
	if err := c.Bind(&newUser); err != nil {
		return utils.HandleError(c, err)
	}

	if err := c.Validate(&newUser); err != nil {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			},
		)
	}

	createdUser, err := ac.AuthService.Register(&newUser)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.WriteResponse(
		c,
		http.StatusCreated,
		createdUser,
		"user successfully created",
	)
}

func (ac AuthController) Login(c echo.Context) error {
	login := struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}
	if err := c.Bind(&login); err != nil {
		return utils.HandleError(c, err)
	}

	if err := c.Validate(&login); err != nil {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			},
		)
	}

	token, err := ac.AuthService.Login(login.Email, login.Password)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.WriteResponse(
		c,
		http.StatusOK,
		struct {
			Token string `json:"token"`
		}{Token: token},
		"login success",
	)
}

func (ac AuthController) ForgetPassword(c echo.Context) error {
	userIdRaw := c.Request().Header.Get("UserId")
	if userIdRaw == "" {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnauthorized,
				Message: "invalid token",
			},
		)
	}

	userId, err := strconv.Atoi(userIdRaw)
	if err != nil {
		return utils.HandleError(c, err)
	}

	newPwdStruct := struct {
		NewPassword string `json:"new_password" validate:"required"`
	}{}
	if err := c.Bind(&newPwdStruct); err != nil {
		return utils.HandleError(c, err)
	}

	if err := c.Validate(&newPwdStruct); err != nil {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			},
		)
	}

	ac.AuthService.ForgetPassword(uint(userId), newPwdStruct.NewPassword)

	return utils.WriteResponse(
		c,
		http.StatusOK,
		nil,
		"update password success",
	)
}

func (ac AuthController) DeleteAccount(c echo.Context) error {
	userIdRaw := c.Request().Header.Get("UserId")
	if userIdRaw == "" {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnauthorized,
				Message: "invalid token",
			},
		)
	}

	userId, err := strconv.Atoi(userIdRaw)
	if err != nil {
		return utils.HandleError(c, err)
	}

	ac.AuthService.DeleteAccount(uint(userId))

	return utils.WriteResponse(
		c,
		http.StatusOK,
		nil,
		"account deletion success",
	)
}
