package routers

import (
	controllers "github.com/IqbalLx/alterra-agmc/controllers/user"
	middlewares "github.com/IqbalLx/alterra-agmc/middlewares/auth"

	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Group, userController controllers.IUserController, authMiddleware middlewares.IAuthMiddleware) *echo.Group {
	userRouter := e.Group("/user")

	userRouter.GET("/me", userController.GetMe, authMiddleware.VerifyToken)
	userRouter.PUT("/me", userController.UpdateMe, authMiddleware.VerifyToken)
	userRouter.GET("", userController.GetUsers, authMiddleware.VerifyToken)

	return userRouter
}
