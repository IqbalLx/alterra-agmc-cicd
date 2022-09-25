package services

import "github.com/IqbalLx/alterra-agmc/entities"

type IUserService interface {
	GetUser(id uint) (entities.User, error)
	GetUsers() ([]entities.User, error)
	UpdateUser(id uint, newUser entities.User) (entities.User, error)
}
