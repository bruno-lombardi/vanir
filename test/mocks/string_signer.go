package mocks

import "github.com/stretchr/testify/mock"

type StringSignerMock struct {
	mock.Mock
}

func (s *StringSignerMock) SignedString(key interface{}) (string, error) {
	args := s.Called(key)
	return args.String(0), args.Error(1)
}
