package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	e "github.com/IqbalLx/alterra-agmc/errors"

	"github.com/labstack/echo/v4"
)

func WriteResponse(
	c echo.Context,
	status int,
	content interface{},
	message string,
) error {
	return c.JSON(status, map[string]interface{}{
		"message": message,
		"content": content,
	})
}

func HandleClientErrorResponse(
	c echo.Context,
	err e.ClientError,
) error {
	return WriteResponse(
		c,
		err.Code,
		nil,
		err.Message,
	)
}

func HandleInternalServerErrorResponse(
	c echo.Context,
	err e.InternalServerError,
) error {
	return WriteResponse(
		c,
		http.StatusInternalServerError,
		nil,
		err.Message,
	)
}

func HandleError(c echo.Context, err error) error {
	var (
		clientError         *e.ClientError
		internalServerError *e.InternalServerError
	)

	switch {
	case errors.As(err, &clientError):
		return HandleClientErrorResponse(c, *clientError)
	case errors.As(err, &internalServerError):
		return HandleInternalServerErrorResponse(c, *internalServerError)
	default:
		log.Fatal(fmt.Errorf(err.Error()))
		return HandleInternalServerErrorResponse(c, e.InternalServerError{Message: "Unknown Internal Server Error"})
	}
}
