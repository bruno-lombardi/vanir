package services_test

import (
	"fmt"
	"testing"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/helpers"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
	"vanir/test/data/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthServiceSuite struct {
	suite.Suite
	authService    *services.AuthServiceImpl
	userRepository *mocks.UserRepositoryMock
	hasher         *mocks.HasherMock
	encrypter      *mocks.EncrypterMock
}

func (sut *AuthServiceSuite) BeforeTest(_, _ string) {
	sut.userRepository = &mocks.UserRepositoryMock{}
	sut.hasher = &mocks.HasherMock{}
	sut.encrypter = &mocks.EncrypterMock{}

	sut.authService = services.NewAuthServiceImpl(sut.userRepository, sut.hasher, sut.encrypter)
}

func (sut *AuthServiceSuite) AfterTest(_, _ string) {
	sut.userRepository.On("FindByEmail", mock.Anything).Unset()
	sut.hasher.On("CompareHashes", mock.Anything, mock.Anything).Unset()
	sut.encrypter.On("CreateToken", mock.Anything).Unset()
}

func (sut *AuthServiceSuite) TestShouldReturnTokenWhenValidCredentials() {
	id := helpers.ID("u")
	sut.userRepository.On("FindByEmail", mock.Anything).Return(
		&repositories.UserEntity{ID: id, Name: "Bruno", Email: "bruno@email.com.br", Password: "123456"},
		nil,
	)
	sut.hasher.On("CompareHashes", mock.Anything, mock.Anything).Return(true)
	sut.encrypter.On("CreateToken", mock.Anything).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", nil)
	authResponse, err := sut.authService.Authenticate(&models.AuthCredentials{
		Email:    "bruno@email.com.br",
		Password: "123456",
	})

	sut.Nil(err)
	sut.Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", authResponse.Token)

	sut.userRepository.AssertCalled(sut.T(), "FindByEmail", "bruno@email.com.br")
	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertCalled(sut.T(), "CompareHashes", "123456", []byte("123456"))
	sut.hasher.AssertExpectations(sut.T())
	sut.encrypter.AssertCalled(sut.T(), "CreateToken", id)
	sut.encrypter.AssertExpectations(sut.T())
}

func (sut *AuthServiceSuite) TestShouldReturnErrorWhenInvalidPassword() {
	sut.userRepository.On("FindByEmail", mock.Anything).Return(
		&repositories.UserEntity{ID: helpers.ID("u"), Name: "Bruno", Email: "bruno@email.com.br", Password: "123456"},
		nil,
	)
	sut.hasher.On("CompareHashes", mock.Anything, mock.Anything).Return(false)
	authResponse, err := sut.authService.Authenticate(&models.AuthCredentials{
		Email:    "bruno@email.com.br",
		Password: "123456",
	})

	sut.Nil(authResponse)
	sut.NotNil(err)
	sut.Equal(&protocols.AppError{
		StatusCode: 401,
		Err:        fmt.Errorf("current password is invalid"),
	}, err)
	sut.userRepository.AssertCalled(sut.T(), "FindByEmail", "bruno@email.com.br")
	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertCalled(sut.T(), "CompareHashes", "123456", []byte("123456"))
	sut.hasher.AssertExpectations(sut.T())
	sut.encrypter.AssertNotCalled(sut.T(), mock.Anything)
	sut.encrypter.AssertExpectations(sut.T())
}

func (sut *AuthServiceSuite) TestShouldReturnErrorWhenInvalidEmail() {
	sut.userRepository.On("FindByEmail", mock.Anything).Return(
		nil,
		fmt.Errorf("user not found"),
	)
	authResponse, err := sut.authService.Authenticate(&models.AuthCredentials{
		Email:    "bruno@email.com.br",
		Password: "123456",
	})

	sut.Nil(authResponse)
	sut.NotNil(err)

	sut.userRepository.AssertCalled(sut.T(), "FindByEmail", "bruno@email.com.br")
	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertNotCalled(sut.T(), "CompareHashes", "123456", []byte("123456"))
	sut.hasher.AssertExpectations(sut.T())
	sut.encrypter.AssertNotCalled(sut.T(), mock.Anything)
	sut.encrypter.AssertExpectations(sut.T())
}

func (sut *AuthServiceSuite) TestShouldReturnErrorWhenTokenCreationFails() {
	id := helpers.ID("u")
	sut.userRepository.On("FindByEmail", mock.Anything).Return(
		&repositories.UserEntity{ID: id, Name: "Bruno", Email: "bruno@email.com.br", Password: "123456"},
		nil,
	)
	sut.hasher.On("CompareHashes", mock.Anything, mock.Anything).Return(true)
	sut.encrypter.On("CreateToken", mock.Anything).Return("", fmt.Errorf("an error occurred creating the token"))
	token, err := sut.authService.Authenticate(&models.AuthCredentials{
		Email:    "bruno@email.com.br",
		Password: "123456",
	})

	sut.Empty(token)
	sut.NotNil(err)

	sut.userRepository.AssertCalled(sut.T(), "FindByEmail", "bruno@email.com.br")
	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertCalled(sut.T(), "CompareHashes", "123456", []byte("123456"))
	sut.hasher.AssertExpectations(sut.T())
	sut.encrypter.AssertCalled(sut.T(), "CreateToken", id)
	sut.encrypter.AssertExpectations(sut.T())
}

func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, &AuthServiceSuite{})
}
