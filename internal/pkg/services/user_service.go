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

func GetUserService() UserService {
	if userService == nil {
		userService = &UserServiceImpl{
			userRepository: repositories.GetUserRepository(),
			hasher:         crypto.GetHasher(),
		}
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
		return &models.User{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}, nil
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

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

}

func (u *UserServiceImpl) Get(ID string) (*models.User, error) {
	user, err := u.userRepository.Get(ID)

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, err
}
