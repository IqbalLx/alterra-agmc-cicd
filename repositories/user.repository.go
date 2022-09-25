package repositories

import "github.com/IqbalLx/alterra-agmc/entities"

type IUserRepository interface {
	CheckExists(id uint) (bool, error)
	CheckExistsByEmail(email string) (bool, error)
	CreateUser(user *entities.User) (entities.User, error)
	GetUser(id uint) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetUsers() ([]entities.User, error)
	UpdateUser(id uint, user *entities.User) (entities.User, error)
	UpdatePassword(id uint, newPassword string) (entities.User, error)
	DeleteUser(id uint) error
}
