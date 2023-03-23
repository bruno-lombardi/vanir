package mocks

import (
	"vanir/internal/pkg/data/models"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (s *UserServiceMock) Create(createUserParams *models.CreateUserParams) (*models.User, error) {
	args := s.Called(createUserParams)
	user, ok := args.Get(0).(*models.User)

	if ok {
		return user, nil
	}

	err, ok := args.Get(1).(error)

	if ok {
		return nil, err
	}

	return nil, nil
}

func (s *UserServiceMock) Update(updateUserParams *models.UpdateUserParams) (*models.User, error) {
	args := s.Called(updateUserParams)
	user, ok := args.Get(0).(*models.User)

	if ok {
		return user, nil
	}

	err, ok := args.Get(1).(error)

	if ok {
		return nil, err
	}

	return nil, nil
}

func (s *UserServiceMock) Get(ID string) (*models.User, error) {
	args := s.Called(ID)
	user, ok := args.Get(0).(*models.User)

	if ok {
		return user, nil
	}

	err, ok := args.Get(1).(error)

	if ok {
		return nil, err
	}

	return nil, nil
}
