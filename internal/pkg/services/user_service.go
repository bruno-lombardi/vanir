package services

import (
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
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
	// TODO: Validate user current password is current

	// TODO: Confirm new password is equal confirmation
	updateUserDTO.NewPassword = crypto.HashAndSalt([]byte(updateUserDTO.NewPassword))
	user, err := u.userRepository.Update(updateUserDTO)

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
