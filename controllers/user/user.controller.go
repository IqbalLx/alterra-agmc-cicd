package controllers

import (
	"net/http"
	"strconv"

	"github.com/IqbalLx/alterra-agmc/entities"
	"github.com/IqbalLx/alterra-agmc/errors"
	services "github.com/IqbalLx/alterra-agmc/services/user"
	"github.com/IqbalLx/alterra-agmc/utils"
	"github.com/labstack/echo/v4"
)

type userController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *userController {
	return &userController{userService}
}

func (uc *userController) GetMe(c echo.Context) error {
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

	user, err := uc.userService.GetUser(uint(userId))
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.WriteResponse(
		c,
		http.StatusOK,
		user,
		"fetch success",
	)
}

func (uc *userController) GetUsers(c echo.Context) error {
	users, err := uc.userService.GetUsers()
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.WriteResponse(
		c,
		http.StatusOK,
		users,
		"fetch success",
	)
}

func (uc *userController) UpdateMe(c echo.Context) error {
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

	newUser := entities.UpdateUserDTO{}
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

	user, err := uc.userService.UpdateUser(uint(userId), entities.User{Username: newUser.Username, Email: newUser.Email})
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.WriteResponse(
		c,
		http.StatusOK,
		user,
		"update success",
	)
}
