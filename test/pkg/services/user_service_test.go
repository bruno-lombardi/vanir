package services_test

import (
	"fmt"
	"testing"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/helpers"
	"vanir/internal/pkg/services"
	"vanir/test/data/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuite struct {
	suite.Suite
	userService    *services.UserServiceImpl
	userRepository *mocks.UserRepositoryMock
	hasher         *mocks.HasherMock
}

func (sut *UserServiceSuite) BeforeTest(_, _ string) {
	sut.userRepository = &mocks.UserRepositoryMock{}
	sut.hasher = &mocks.HasherMock{}

	sut.userService = services.NewUserServiceImpl(sut.userRepository, sut.hasher)
}

func (sut *UserServiceSuite) AfterTest(_, _ string) {
	sut.userRepository.On("Update", mock.Anything).Unset()
	sut.userRepository.On("Create", mock.Anything).Unset()
	sut.userRepository.On("Get", mock.Anything).Unset()
	sut.hasher.On("CompareHashes", mock.Anything).Unset()
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

func (sut *UserServiceSuite) TestShouldReturnErrorIfUserRepositoryFails() {
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

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, &UserServiceSuite{})
}
