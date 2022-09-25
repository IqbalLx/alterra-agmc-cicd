package middlewares

import "github.com/labstack/echo/v4"

type IAuthMiddleware interface {
	VerifyToken(next echo.HandlerFunc) echo.HandlerFunc
}
