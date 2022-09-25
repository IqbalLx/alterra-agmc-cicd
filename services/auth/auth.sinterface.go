package services

import "github.com/IqbalLx/alterra-agmc/entities"

type IAuthService interface {
	Register(user *entities.User) (entities.User, error)
	Login(email, password string) (string, error)
	VerifyToken(authToken string) (entities.Token, error)
	ForgetPassword(id uint, newPassword string) (entities.User, error)
	DeleteAccount(id uint) error
}
