package mocks

import (
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func eitherUserEntityOrError(args mock.Arguments) (*repositories.UserEntity, error) {
	user, ok := args.Get(0).(*repositories.UserEntity)

	if ok {
		return user, nil
	}

	return nil, args.Error(1)
}

func (r *UserRepositoryMock) Get(ID string) (*repositories.UserEntity, error) {
	args := r.Called(ID)
	return eitherUserEntityOrError(args)
}

func (r *UserRepositoryMock) FindByEmail(email string) (*repositories.UserEntity, error) {
	args := r.Called(email)
	return eitherUserEntityOrError(args)
}

func (r *UserRepositoryMock) Create(createUserParams *models.CreateUserParams) (*repositories.UserEntity, error) {
	args := r.Called(createUserParams)
	return eitherUserEntityOrError(args)
}

func (r *UserRepositoryMock) Update(updateUserParams *models.UpdateUserParams) (*repositories.UserEntity, error) {
	args := r.Called(updateUserParams)
	return eitherUserEntityOrError(args)
}
