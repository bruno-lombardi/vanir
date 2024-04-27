package controllers

import (
	"fmt"
	"net/http"
	"testing"
	"vanir/internal/pkg/config"
	"vanir/internal/pkg/data/db"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/helpers"
	"vanir/internal/pkg/protocols"
	data_test "vanir/test/data"
	mocks "vanir/test/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ControllerTestCase data_test.ControllerTestCase

func TestGetUserControllerTestCases(t *testing.T) {
	config.Setup()
	db.SetupDB()

	userServiceMock := &mocks.UserServiceMock{}

	controllerTestCases := []ControllerTestCase{
		{
			Name: "Should return user when service returns the user with valid id",
			WhenRequest: &protocols.HttpRequest{
				PathParams: map[string]string{
					"id": "123",
				},
			},
			BeforeTest: func() error {
				user := &models.User{ID: helpers.ID("u"), Email: "bruno@email.com", Password: "654321", Name: "Bruno Lombardi"}
				userServiceMock.On("Get", mock.Anything).Return(user, nil)
				return nil
			},
			ExpectResponse: func(t *testing.T, response *protocols.HttpResponse) error {
				user, ok := response.Body.(*models.User)
				require.True(t, ok)
				require.NotNil(t, user.ID)
				require.Equal(t, http.StatusOK, response.StatusCode)
				require.Equal(t, "bruno@email.com", user.Email)
				return nil
			},
			AfterTest: func() error {
				userServiceMock.AssertCalled(t, "Get", "123")
				userServiceMock.AssertExpectations(t)
				userServiceMock.On("Get", mock.Anything).Unset()
				return nil
			},
		},
		{
			Name: "Should not return user when user service returns error",
			WhenRequest: &protocols.HttpRequest{
				PathParams: map[string]string{
					"id": "123",
				},
			},
			BeforeTest: func() error {
				userServiceMock.On("Get", mock.Anything).Return(nil, fmt.Errorf("user service threw an error"))
				return nil
			},
			ExpectResponse: func(t *testing.T, response *protocols.HttpResponse) error {
				require.Equal(t, http.StatusInternalServerError, response.StatusCode)
				return nil
			},
			AfterTest: func() error {
				userServiceMock.AssertCalled(t, "Get", "123")
				userServiceMock.AssertExpectations(t)
				userServiceMock.On("Get", mock.Anything).Unset()
				return nil
			},
		},
	}

	for _, testCase := range controllerTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			require.NoError(t, testCase.BeforeTest())

			updateUserController := NewGetUserController(
				userServiceMock,
			)
			request, ok := testCase.WhenRequest.(*protocols.HttpRequest)
			require.True(t, ok)

			response, _ := updateUserController.Handle(request)

			require.NoError(t, testCase.ExpectResponse(t, response))
			require.NoError(t, testCase.AfterTest())
		})
	}
}
