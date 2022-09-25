package controllers

import (
	"github.com/labstack/echo/v4"
)

type IAuthController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	ForgetPassword(e echo.Context) error
	DeleteAccount(e echo.Context) error
}
