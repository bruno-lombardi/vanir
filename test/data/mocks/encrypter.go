package mocks

import "github.com/stretchr/testify/mock"

type EncrypterMock struct {
	mock.Mock
}

func (e *EncrypterMock) CreateToken(subject string) (token string, err error) {
	args := e.Called(subject)
	return args.String(0), args.Error(1)
}

func (e *EncrypterMock) ValidateToken(tokenString string) (valid bool, subject string) {
	args := e.Called(tokenString)
	return args.Bool(0), args.String(1)
}
