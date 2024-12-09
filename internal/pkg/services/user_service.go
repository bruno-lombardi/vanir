package services

import (
	"fmt"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/protocols"
)

type UserService interface {
	Create(createUserParams *models.CreateUserParams) (*models.User, error)
	Get(ID string) (*models.User, error)
	Update(updateUserParams *models.UpdateUserParams) (*models.User, error)
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
	hasher         crypto.Hasher
}

var userService *UserServiceImpl

func GetUserService(userRepository repositories.UserRepository, hasher crypto.Hasher) UserService {
	if userService == nil {
		userService = NewUserServiceImpl(userRepository, hasher)
	}
	return userService
}

func NewUserServiceImpl(userRepository repositories.UserRepository, hasher crypto.Hasher) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
		hasher:         hasher,
	}
}

func (u *UserServiceImpl) Create(createUserParams *models.CreateUserParams) (*models.User, error) {
	createUserParams.Password = u.hasher.HashAndSalt([]byte(createUserParams.Password))
	user, err := u.userRepository.Create(createUserParams)

	if err != nil {
		return nil, err
	} else {
		return mapUserToModel(user), nil
	}

}

func (u *UserServiceImpl) Update(updateUserParams *models.UpdateUserParams) (*models.User, error) {
	user, err := u.userRepository.Get(updateUserParams.ID)
	if err != nil {
		return nil, err
	}

	isCompareSuccessful := u.hasher.CompareHashes(user.Password, []byte(updateUserParams.CurrentPassword))

	if !isCompareSuccessful {
		return nil, &protocols.AppError{
			StatusCode: 401,
			Err:        fmt.Errorf("current password is invalid"),
		}
	}

	updateUserParams.NewPassword = u.hasher.HashAndSalt([]byte(updateUserParams.NewPassword))
	user, err = u.userRepository.Update(updateUserParams)
	if err != nil {
		return nil, err
	}

	return mapUserToModel(user), nil

}

func (u *UserServiceImpl) Get(ID string) (*models.User, error) {
	user, err := u.userRepository.Get(ID)

	return mapUserToModel(user), err
}

func mapUserToModel(user *repositories.UserEntity) *models.User {
	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}
