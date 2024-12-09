package middlewares

import (
	"fmt"
	"net/http"
	"testing"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/test/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthMiddlewareSuite struct {
	suite.Suite
	authMiddleware *AuthenticatedMiddleware
	encrypter      *mocks.EncrypterMock
	userService    *mocks.UserServiceMock
}

func (sut *AuthMiddlewareSuite) BeforeTest(_, _ string) {
	sut.userService = &mocks.UserServiceMock{}
	sut.encrypter = &mocks.EncrypterMock{}

	sut.authMiddleware = NewAuthenticatedMiddleware(sut.encrypter, sut.userService)
}

func (sut *AuthMiddlewareSuite) AfterTest(_, _ string) {
	sut.userService.On("Get", mock.Anything).Unset()
	sut.encrypter.On("ValidateToken", mock.Anything).Unset()
}

func (sut *AuthMiddlewareSuite) TestSmokeTest() {
	sut.NotNil(GetAuthenticatedMiddleware(&mocks.EncrypterMock{}, &mocks.UserServiceMock{}))
}

func (sut *AuthMiddlewareSuite) TestShouldNotReturnErrorIfAuthorizationTokenIsValid() {
	sut.encrypter.On("ValidateToken", mock.Anything).Return(true, "u_uhu13oasjdf")
	sut.userService.On("Get", mock.Anything).Return(&models.User{ID: "u_uhu13oasjdf"}, nil)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	req := &protocols.HttpRequest{
		HttpReq: &http.Request{
			Header: map[string][]string{
				"Content-Type":  {"application/json"},
				"Authorization": {fmt.Sprintf("Bearer %s", token)},
			},
		},
	}
	err := sut.authMiddleware.Handle(req)

	sut.Nil(err)
	sut.NotNil(req.AuthenticatedUser)
	sut.Equal("u_uhu13oasjdf", req.AuthenticatedUser.ID)

	sut.encrypter.AssertCalled(sut.T(), "ValidateToken", token)
	sut.userService.AssertExpectations(sut.T())
}

func (sut *AuthMiddlewareSuite) TestShouldReturnErrorIfAuthorizationTokenIsInvalidOrExpired() {
	sut.encrypter.On("ValidateToken", mock.Anything).Return(false, "u_uhu13oasjdf")
	sut.userService.On("Get", mock.Anything).Return(&models.User{ID: "u_uhu13oasjdf"}, nil)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	req := &protocols.HttpRequest{
		HttpReq: &http.Request{
			Header: map[string][]string{
				"Content-Type":  {"application/json"},
				"Authorization": {fmt.Sprintf("Bearer %s", token)},
			},
		},
	}
	err := sut.authMiddleware.Handle(req)

	sut.NotNil(err)
	sut.Equal(http.StatusUnauthorized, err.(*protocols.AppError).StatusCode)

	sut.encrypter.AssertCalled(sut.T(), "ValidateToken", token)
	sut.userService.AssertNotCalled(sut.T(), "Get", mock.Anything)
}

func (sut *AuthMiddlewareSuite) TestShouldReturnErrorIfIfNoAuthorizationHeaderIsSent() {
	req := &protocols.HttpRequest{
		HttpReq: &http.Request{
			Header: map[string][]string{
				"Content-Type":  {"application/json"},
				"Authorization": {""},
			},
		},
	}
	err := sut.authMiddleware.Handle(req)

	sut.NotNil(err)
	sut.Equal(http.StatusUnauthorized, err.(*protocols.AppError).StatusCode)

	sut.encrypter.AssertNotCalled(sut.T(), "ValidateToken", "")
}

func (sut *AuthMiddlewareSuite) TestShouldReturnErrorIfUserIsNotFoundAsSubject() {
	sut.encrypter.On("ValidateToken", mock.Anything).Return(true, "u_uhu13oasjdf")
	sut.userService.On("Get", mock.Anything).Return(nil, fmt.Errorf("user was not found"))
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	req := &protocols.HttpRequest{
		HttpReq: &http.Request{
			Header: map[string][]string{
				"Content-Type":  {"application/json"},
				"Authorization": {fmt.Sprintf("Bearer %s", token)},
			},
		},
	}
	err := sut.authMiddleware.Handle(req)

	sut.NotNil(err)
	sut.Equal(http.StatusUnauthorized, err.(*protocols.AppError).StatusCode)

	sut.encrypter.AssertCalled(sut.T(), "ValidateToken", token)
	sut.userService.AssertCalled(sut.T(), "Get", mock.Anything)
}

func TestAuthMiddlewareSuite(t *testing.T) {
	suite.Run(t, &AuthMiddlewareSuite{})
}
