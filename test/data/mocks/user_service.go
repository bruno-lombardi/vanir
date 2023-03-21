package mocks

import (
	"vanir/internal/pkg/data/models"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (s *UserServiceMock) Create(createUserDTO *models.CreateUserDTO) (*models.User, error) {
	args := s.Called(createUserDTO)
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

func (s *UserServiceMock) Update(updateUserDTO *models.UpdateUserDTO) (*models.User, error) {
	args := s.Called(updateUserDTO)
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
