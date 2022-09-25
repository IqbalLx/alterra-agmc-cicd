package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	services "github.com/IqbalLx/alterra-agmc/services/auth"
	"github.com/IqbalLx/alterra-agmc/utils"
	"github.com/labstack/echo/v4"
)

type authMiddleware struct {
	authService services.IAuthService
}

func NewAuthMiddleware(authService services.IAuthService) *authMiddleware {
	return &authMiddleware{authService}
}

func (am *authMiddleware) VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return utils.WriteResponse(
				c,
				http.StatusUnauthorized,
				nil,
				"invalid token",
			)
		}

		splittedToken := strings.Fields(token)
		if splittedToken[0] != "Bearer" {
			return utils.WriteResponse(
				c,
				http.StatusUnauthorized,
				nil,
				"invalid token",
			)
		}

		authToken := splittedToken[1]
		tokenEntity, err := am.authService.VerifyToken(authToken)
		if err != nil {
			return utils.HandleError(c, err)
		}

		// still manually map each key
		c.Request().Header.Set("UserId", strconv.Itoa(int(tokenEntity.UserId)))

		return next(c)
	}
}
