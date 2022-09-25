package routers

import (
	controllers "github.com/IqbalLx/alterra-agmc/controllers/auth"
	middlewares "github.com/IqbalLx/alterra-agmc/middlewares/auth"

	"github.com/labstack/echo/v4"
)

func NewAuthRouter(e *echo.Group, authController controllers.IAuthController, authMiddleware middlewares.IAuthMiddleware) *echo.Group {
	authRouter := e.Group("/auth")

	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.Login)
	authRouter.PUT("", authController.ForgetPassword, authMiddleware.VerifyToken)
	authRouter.DELETE("", authController.DeleteAccount, authMiddleware.VerifyToken)

	return authRouter
}
