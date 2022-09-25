package mysqlDataSource

import (
	m "github.com/IqbalLx/alterra-agmc/dataSources/mysql/schema"
	"github.com/IqbalLx/alterra-agmc/entities"
	e "github.com/IqbalLx/alterra-agmc/errors"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	gorm *gorm.DB
}

func (msql *MysqlUserRepository) CheckExists(id uint) (bool, error) {
	var count int64
	if res := msql.gorm.Table("users").Where("id = ?", id).Count(&count); res.Error != nil {
		return false, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	return count > 0, nil
}

func (msql *MysqlUserRepository) CheckExistsByEmail(email string) (bool, error) {
	var count int64
	if res := msql.gorm.Table("users").Where("email = ?", email).Count(&count); res.Error != nil {
		return false, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	return count > 0, nil
}

func (msql *MysqlUserRepository) CreateUser(user *entities.User) (entities.User, error) {
	userDB := m.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if res := msql.gorm.Table("users").Create(&userDB); res.Error != nil {
		return entities.User{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	userEntity := entities.User{
		Id:       userDB.ID,
		Username: userDB.Username,
		Email:    userDB.Email, Password: userDB.Password,
	}

	return userEntity, nil
}

func (msql *MysqlUserRepository) GetUser(id uint) (entities.User, error) {
	user := m.User{}
	if res := msql.gorm.First(&user, id); res.Error != nil {
		return entities.User{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	userEntity := entities.User{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	return userEntity, nil
}

func (msql *MysqlUserRepository) GetUserByEmail(email string) (entities.User, error) {
	user := m.User{}
	if res := msql.gorm.Table("users").Where("email = ?", email).First(&user); res.Error != nil {
		return entities.User{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	userEntity := entities.User{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	return userEntity, nil
}

func (msql *MysqlUserRepository) GetUsers() ([]entities.User, error) {
	users := []m.User{}
	if res := msql.gorm.Find(&users); res.Error != nil {
		return []entities.User{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	usersEntity := []entities.User{}
	for idx := 0; idx < len(users); idx++ {
		usersEntity = append(usersEntity, entities.User{
			Id:       users[idx].ID,
			Username: users[idx].Username,
			Email:    users[idx].Email,
			Password: users[idx].Password,
		})
	}

	return usersEntity, nil
}

func (msql *MysqlUserRepository) UpdateUser(id uint, user *entities.User) (entities.User, error) {
	userDB := &m.User{}
	if res := msql.gorm.Table("users").Where("id = ?", id).First(userDB).Updates(m.User{Username: user.Username, Email: user.Email}); res.Error != nil {
		return entities.User{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	userEntity := entities.User{
		Id:       userDB.ID,
		Username: userDB.Username,
		Email:    userDB.Email,
		Password: userDB.Password,
	}

	return userEntity, nil
}

func (msql *MysqlUserRepository) UpdatePassword(id uint, newPassword string) (entities.User, error) {
	userDB := &m.User{}
	if res := msql.gorm.Table("users").Where("id = ?", id).First(userDB).Update("password", newPassword); res.Error != nil {
		return entities.User{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	userEntity := entities.User{
		Id:       userDB.ID,
		Username: userDB.Username,
		Email:    userDB.Email,
		Password: userDB.Password,
	}

	return userEntity, nil
}

func (msql *MysqlUserRepository) DeleteUser(id uint) error {
	if res := msql.gorm.Unscoped().Delete(&m.User{}, id); res.Error != nil {
		return &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	return nil
}

func NewMysqlUserRepository(gorm *gorm.DB) *MysqlUserRepository {
	return &MysqlUserRepository{gorm}
}
