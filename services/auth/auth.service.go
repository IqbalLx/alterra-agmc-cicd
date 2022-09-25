package services

import (
	"fmt"
	"net/http"

	"github.com/IqbalLx/alterra-agmc/entities"
	e "github.com/IqbalLx/alterra-agmc/errors"
	"github.com/IqbalLx/alterra-agmc/repositories"
	"github.com/IqbalLx/alterra-agmc/utils"
)

type AuthService struct {
	userRepository repositories.IUserRepository
	hashUtil       utils.IHashUtils
	jwtUtil        utils.IJWTUtil
}

func NewAuthService(userRepository repositories.IUserRepository, hashUtil utils.IHashUtils, jwtUtil utils.IJWTUtil) *AuthService {
	return &AuthService{userRepository, hashUtil, jwtUtil}
}

func (as *AuthService) Register(user *entities.User) (entities.User, error) {
	isUserExists, err := as.userRepository.CheckExistsByEmail(user.Email)
	if err != nil {
		return entities.User{}, err
	}

	if isUserExists {
		return entities.User{}, &e.ClientError{
			Code:    http.StatusConflict,
			Message: fmt.Sprintf("user with email %s is exists", user.Email),
		}
	}

	hashedPassword, err := as.hashUtil.Hash(user.Password)
	if err != nil {
		return entities.User{}, &e.InternalServerError{
			Message: "Internal Server Error",
		}
	}

	user.Password = hashedPassword

	createdUser, err := as.userRepository.CreateUser(user)
	if err != nil {
		return entities.User{}, &e.ClientError{
			Code:    http.StatusConflict,
			Message: fmt.Sprintf("user with email %s is exists", user.Email),
		}
	}

	return createdUser, nil
}

func (as *AuthService) Login(email, password string) (string, error) {
	isUserExists, err := as.userRepository.CheckExistsByEmail(email)
	if err != nil {
		return "", err
	}

	if !isUserExists {
		return "", &e.ClientError{
			Code:    http.StatusUnauthorized,
			Message: fmt.Sprintf("user with email %s is not exists", email),
		}
	}

	user, err := as.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	isPasswordValid := as.hashUtil.Compare(password, user.Password)
	if !isPasswordValid {
		return "", &e.ClientError{
			Code:    http.StatusUnauthorized,
			Message: "password mismatched",
		}
	}

	token, err := as.jwtUtil.Generate(user.Id)
	if err != nil {
		return "", &e.InternalServerError{
			Message: "Internal Server Error",
		}
	}

	return token, nil
}

func (as *AuthService) VerifyToken(authToken string) (entities.Token, error) {
	token, err := as.jwtUtil.Validate(authToken)
	if err != nil {
		return entities.Token{}, &e.ClientError{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}
	}

	isUserExists, err := as.userRepository.CheckExists(token.UserId)
	if err != nil {
		return entities.Token{}, err
	}
	if !isUserExists {
		return entities.Token{}, &e.ClientError{
			Code:    http.StatusUnauthorized,
			Message: "user not exists",
		}
	}

	return token, nil
}

func (as *AuthService) ForgetPassword(id uint, newPassword string) (entities.User, error) {
	isUserExists, err := as.userRepository.CheckExists(id)
	if err != nil {
		return entities.User{}, err
	}
	if !isUserExists {
		return entities.User{}, &e.ClientError{
			Code:    http.StatusUnauthorized,
			Message: "user not exists",
		}
	}

	hashedPassword, err := as.hashUtil.Hash(newPassword)
	if err != nil {
		return entities.User{}, &e.InternalServerError{
			Message: "internal server error",
		}
	}

	updatedUser, err := as.userRepository.UpdatePassword(id, hashedPassword)
	if err != nil {
		return entities.User{}, err
	}

	return updatedUser, nil
}

func (as *AuthService) DeleteAccount(id uint) error {
	isUserExists, err := as.userRepository.CheckExists(id)
	if err != nil {
		return err
	}
	if !isUserExists {
		return &e.ClientError{
			Code:    http.StatusUnauthorized,
			Message: "user not exists",
		}
	}

	if err := as.userRepository.DeleteUser(id); err != nil {
		return err
	}

	return nil
}
