package routers

import (
	controllers "github.com/IqbalLx/alterra-agmc/controllers/book"
	middlewares "github.com/IqbalLx/alterra-agmc/middlewares/auth"

	"github.com/labstack/echo/v4"
)

func NewBookRouter(e *echo.Group, bookController controllers.IBookController, authMiddleware middlewares.IAuthMiddleware) *echo.Group {
	bookRouter := e.Group("/book")

	bookRouter.POST("", bookController.CreateBook, authMiddleware.VerifyToken)
	bookRouter.GET("/:id", bookController.GetBook, authMiddleware.VerifyToken)
	bookRouter.GET("", bookController.GetBooks, authMiddleware.VerifyToken)
	bookRouter.PUT("/:id", bookController.UpdateBook, authMiddleware.VerifyToken)
	bookRouter.DELETE("/:id", bookController.DeleteBook, authMiddleware.VerifyToken)

	return bookRouter
}
