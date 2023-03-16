package services

import (
	"fmt"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/protocols"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

var userService *UserService

func GetUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			userRepository: repositories.GetUserRepository(),
		}
	}
	return userService
}

func (u *UserService) Create(createUserDTO *models.CreateUserDTO) (*models.User, error) {
	createUserDTO.Password = crypto.HashAndSalt([]byte(createUserDTO.Password))
	user, err := u.userRepository.Create(createUserDTO)

	if err != nil {
		return nil, err
	} else {
		return &models.User{
			ID:       user.ID,
			Email:    user.Email,
			Name:     user.Name,
			Password: user.Password,
		}, nil
	}

}

func (u *UserService) Update(updateUserDTO *models.UpdateUserDTO) (*models.User, error) {
	user, err := u.userRepository.Get(updateUserDTO.ID)
	if err != nil {
		return nil, err
	}

	isCompareSuccessful := crypto.CompareHashes(user.Password, []byte(updateUserDTO.CurrentPassword))

	if !isCompareSuccessful {
		return nil, &protocols.AppError{
			StatusCode: 400,
			Err:        fmt.Errorf("current password is invalid"),
		}
	}

	updateUserDTO.NewPassword = crypto.HashAndSalt([]byte(updateUserDTO.NewPassword))
	user, err = u.userRepository.Update(updateUserDTO)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}, nil

}
