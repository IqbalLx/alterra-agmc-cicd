package services

import (
	"github.com/IqbalLx/alterra-agmc/entities"
	"github.com/IqbalLx/alterra-agmc/repositories"
)

type userService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *userService {
	return &userService{userRepository}
}

func (us *userService) GetUser(id uint) (entities.User, error) {
	return us.userRepository.GetUser(id)

}

func (us *userService) GetUsers() ([]entities.User, error) {
	return us.userRepository.GetUsers()
}

func (us *userService) UpdateUser(id uint, newUser entities.User) (entities.User, error) {
	return us.userRepository.UpdateUser(id, &newUser)
}
