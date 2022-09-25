package controllers

import "github.com/labstack/echo/v4"

type IUserController interface {
	GetMe(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateMe(c echo.Context) error
}
