package mocks

import "github.com/stretchr/testify/mock"

type HasherMock struct {
	mock.Mock
}

func (h *HasherMock) HashAndSalt(value []byte) string {
	args := h.Called(value)
	return args.String(0)
}

func (h *HasherMock) CompareHashes(hashed string, plain []byte) bool {
	args := h.Called(hashed, plain)
	return args.Bool(0)
}
