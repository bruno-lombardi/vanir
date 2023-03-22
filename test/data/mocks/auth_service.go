package mocks

import (
	"vanir/internal/pkg/data/models"

	"github.com/stretchr/testify/mock"
)

type AuthServiceMock struct {
	mock.Mock
}

func (s *AuthServiceMock) Authenticate(authCredentialsDTO *models.AuthCredentialsDTO) (token string, err error) {
	args := s.Called(authCredentialsDTO)
	token = args.String(0)
	err = args.Error(1)

	return token, err
}
