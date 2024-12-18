package services

import (
	"fmt"
	"testing"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/helpers"
	"vanir/test/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuite struct {
	suite.Suite
	userService    *UserServiceImpl
	userRepository *mocks.UserRepositoryMock
	hasher         *mocks.HasherMock
}

func (sut *UserServiceSuite) BeforeTest(_, _ string) {
	sut.userRepository = &mocks.UserRepositoryMock{}
	sut.hasher = &mocks.HasherMock{}

	sut.userService = NewUserServiceImpl(sut.userRepository, sut.hasher)
}

func (sut *UserServiceSuite) AfterTest(_, _ string) {
	sut.userRepository.On("Update", mock.Anything).Unset()
	sut.userRepository.On("Create", mock.Anything).Unset()
	sut.userRepository.On("Get", mock.Anything).Unset()
	sut.hasher.On("CompareHashes", mock.Anything).Unset()
}

func (sut *UserServiceSuite) TestSmokeTest() {
	sut.NotNil(GetUserService(&mocks.UserRepositoryMock{}, &mocks.HasherMock{}))
}

func (sut *UserServiceSuite) TestShouldCreateUserWhenValidData() {
	id := helpers.ID("u")
	sut.hasher.On("HashAndSalt", mock.Anything).Return("$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O")
	sut.userRepository.On("Create", mock.Anything).Return(
		&repositories.UserEntity{ID: id, Name: "Bruno", Email: "bruno@email.com.br", Password: "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O"},
		nil,
	)

	user, err := sut.userService.Create(&models.CreateUserParams{
		Name:                 "Bruno",
		Email:                "bruno@email.com.br",
		Password:             "123456",
		PasswordConfirmation: "123456",
	})

	sut.Nil(err)
	sut.Equal(user.ID, id)
	sut.Equal("$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O", user.Password)

	sut.userRepository.AssertCalled(sut.T(), "Create", &models.CreateUserParams{
		Name:                 "Bruno",
		Email:                "bruno@email.com.br",
		Password:             "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O",
		PasswordConfirmation: "123456",
	})
	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertCalled(sut.T(), "HashAndSalt", []byte("123456"))
	sut.hasher.AssertExpectations(sut.T())
}

func (sut *UserServiceSuite) TestShouldReturnErrorIfUserRepositoryFailsOnCreateUser() {
	sut.hasher.On("HashAndSalt", mock.Anything).Return("$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O")
	sut.userRepository.On("Create", mock.Anything).Return(
		nil,
		fmt.Errorf("user repository threw error"),
	)

	user, err := sut.userService.Create(&models.CreateUserParams{
		Name:                 "Bruno",
		Email:                "bruno@email.com.br",
		Password:             "123456",
		PasswordConfirmation: "123456",
	})

	sut.NotNil(err)
	sut.Nil(user)

	sut.userRepository.AssertCalled(sut.T(), "Create", &models.CreateUserParams{
		Name:                 "Bruno",
		Email:                "bruno@email.com.br",
		Password:             "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O",
		PasswordConfirmation: "123456",
	})
	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertCalled(sut.T(), "HashAndSalt", []byte("123456"))
	sut.hasher.AssertExpectations(sut.T())
}

func (sut *UserServiceSuite) TestShouldGetUserSuccessful() {
	id := helpers.ID("u")
	sut.userRepository.On("Get", mock.Anything).Return(
		&repositories.UserEntity{ID: id, Name: "Bruno", Email: "bruno@email.com.br", Password: "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O"},
		nil,
	)

	user, err := sut.userService.Get(id)

	sut.Nil(err)
	sut.Equal(user.ID, id)

	sut.userRepository.AssertExpectations(sut.T())
}

func (sut *UserServiceSuite) TestShouldUpdateUserWhenValidData() {
	id := helpers.ID("u")
	sut.hasher.On("CompareHashes", mock.Anything, mock.Anything).Return(true)
	sut.hasher.On("HashAndSalt", mock.Anything).Return("$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O")
	sut.userRepository.On("Get", mock.Anything).Return(
		&repositories.UserEntity{ID: id, Name: "Bruno", Email: "bruno@email.com.br", Password: "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O"},
		nil,
	)
	sut.userRepository.On("Update", mock.Anything).Return(
		&repositories.UserEntity{ID: id, Name: "Bruno", Email: "bruno@email.com.br", Password: "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O"},
		nil,
	)

	user, err := sut.userService.Update(&models.UpdateUserParams{
		Name:                    "Bruno",
		Email:                   "bruno@email.com.br",
		CurrentPassword:         "123456",
		NewPassword:             "654321",
		NewPasswordConfirmation: "654321",
	})

	sut.Nil(err)
	sut.Equal(user.ID, id)
	sut.Equal("$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O", user.Password)

	sut.userRepository.AssertCalled(sut.T(), "Update", &models.UpdateUserParams{
		Name:                    "Bruno",
		Email:                   "bruno@email.com.br",
		CurrentPassword:         "123456",
		NewPassword:             "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O",
		NewPasswordConfirmation: "654321",
	})
	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertCalled(sut.T(), "HashAndSalt", []byte("654321"))
	sut.hasher.AssertExpectations(sut.T())
}

func (sut *UserServiceSuite) TestShouldFailUpdateUserWhenRepositoryFails() {
	sut.userRepository.On("Get", mock.Anything).Return(
		nil,
		fmt.Errorf("error when get user"),
	)

	user, err := sut.userService.Update(&models.UpdateUserParams{
		Name:                    "Bruno",
		Email:                   "bruno@email.com.br",
		CurrentPassword:         "123456",
		NewPassword:             "654321",
		NewPasswordConfirmation: "654321",
	})

	sut.NotNil(err)
	sut.Nil(user)

	sut.userRepository.AssertExpectations(sut.T())
}

func (sut *UserServiceSuite) TestShouldFailUpdateUserWhenPasswordIsInvalid() {
	id := helpers.ID("u")
	sut.hasher.On("CompareHashes", mock.Anything, mock.Anything).Return(false)
	sut.userRepository.On("Get", mock.Anything).Return(
		&repositories.UserEntity{ID: id, Name: "Bruno", Email: "bruno@email.com.br", Password: "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O"},
		nil,
	)

	user, err := sut.userService.Update(&models.UpdateUserParams{
		Name:                    "Bruno",
		Email:                   "bruno@email.com.br",
		CurrentPassword:         "123456",
		NewPassword:             "654321",
		NewPasswordConfirmation: "654321",
	})

	sut.NotNil(err)
	sut.Nil(user)
	sut.ErrorContains(err, "current password is invalid")

	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertExpectations(sut.T())
}

func (sut *UserServiceSuite) TestShouldFailUpdateUserWhenRepositoryUpdateFails() {
	id := helpers.ID("u")
	sut.hasher.On("CompareHashes", mock.Anything, mock.Anything).Return(true)
	sut.hasher.On("HashAndSalt", mock.Anything).Return("$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O")
	sut.userRepository.On("Get", mock.Anything).Return(
		&repositories.UserEntity{ID: id, Name: "Bruno", Email: "bruno@email.com.br", Password: "$2a$12$rWgChwk828BWU3bRWEx6M.WlLRNisVPsL47hH7ilYcaE4NxNFQw/O"},
		nil,
	)
	sut.userRepository.On("Update", mock.Anything).Return(
		nil,
		fmt.Errorf("error on update"),
	)

	user, err := sut.userService.Update(&models.UpdateUserParams{
		Name:                    "Bruno",
		Email:                   "bruno@email.com.br",
		CurrentPassword:         "123456",
		NewPassword:             "654321",
		NewPasswordConfirmation: "654321",
	})

	sut.NotNil(err)
	sut.Nil(user)
	sut.ErrorContains(err, "error on update")

	sut.userRepository.AssertExpectations(sut.T())
	sut.hasher.AssertExpectations(sut.T())
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, &UserServiceSuite{})
}
