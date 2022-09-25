package main

import (
	"github.com/IqbalLx/alterra-agmc/config"
	authController "github.com/IqbalLx/alterra-agmc/controllers/auth"
	bookController "github.com/IqbalLx/alterra-agmc/controllers/book"
	userController "github.com/IqbalLx/alterra-agmc/controllers/user"
	mysqlDataSource "github.com/IqbalLx/alterra-agmc/dataSources/mysql"
	middlewares "github.com/IqbalLx/alterra-agmc/middlewares/auth"
	"github.com/IqbalLx/alterra-agmc/routers"
	authService "github.com/IqbalLx/alterra-agmc/services/auth"
	bookService "github.com/IqbalLx/alterra-agmc/services/book"
	userService "github.com/IqbalLx/alterra-agmc/services/user"
	"github.com/IqbalLx/alterra-agmc/utils"
	"github.com/labstack/echo/v4"
)

func InitializeServer(api *echo.Group) *echo.Group {
	config := config.GetConfig()
	db := mysqlDataSource.InitDB(config)

	// repositories
	mysqlUserRepo := mysqlDataSource.NewMysqlUserRepository(db)
	mysqlBookRepo := mysqlDataSource.NewMysqlBookRepository(db)

	// utils
	hashUtil := utils.NewBcryptHashUtils()
	jwtUtil := utils.NewGoJWTUtil(config.Auth.JWTSecret)

	// services
	authService := authService.NewAuthService(mysqlUserRepo, hashUtil, jwtUtil)
	userService := userService.NewUserService(mysqlUserRepo)
	bookService := bookService.NewBookService(mysqlBookRepo)

	// middlewares
	authMiddleware := middlewares.NewAuthMiddleware(authService)

	// controller
	authController := authController.NewAuthController(authService)
	userController := userController.NewUserController(userService)
	bookController := bookController.NewBookController(bookService)

	// router
	routers.NewAuthRouter(api, authController, authMiddleware)
	routers.NewUserRouter(api, userController, authMiddleware)
	routers.NewBookRouter(api, bookController, authMiddleware)

	return api
}
