package user_controllers_test

import (
	"fmt"
	"net/http"
	"testing"
	controllers "vanir/internal/app/presentation/controllers/users"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/helpers"
	"vanir/internal/pkg/protocols"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRunUpdateUserControllerTestCases(t *testing.T) {
	controllerTestCases := []ControllerTestCase{
		{
			Name: "Should update user with valid data",
			WhenRequest: &protocols.HttpRequest{
				Body: &models.UpdateUserParams{
					Name:                    "Bruno Lombardi",
					Email:                   "bruno@email.com",
					CurrentPassword:         "123456",
					NewPassword:             "654321",
					NewPasswordConfirmation: "654321",
				},
			},
			BeforeTest: func() error {
				user := &models.User{ID: helpers.ID("u"), Email: "bruno@email.com", Password: "654321", Name: "Bruno Lombardi"}
				userServiceMock.On("Update", mock.Anything).Return(user, nil)
				return nil
			},
			ExpectResponse: func(t *testing.T, response *protocols.HttpResponse) error {
				user, ok := response.Body.(*models.User)
				assert.True(t, ok)
				assert.NotNil(t, user.ID)
				assert.Equal(t, http.StatusOK, response.StatusCode)
				assert.Equal(t, "bruno@email.com", user.Email)
				return nil
			},
			AfterTest: func() error {
				userServiceMock.AssertCalled(t, "Update", &models.UpdateUserParams{
					Name:                    "Bruno Lombardi",
					Email:                   "bruno@email.com",
					CurrentPassword:         "123456",
					NewPassword:             "654321",
					NewPasswordConfirmation: "654321",
				})
				userServiceMock.AssertExpectations(t)
				userServiceMock.On("Update", mock.Anything).Unset()
				return nil
			},
		},
		{
			Name: "Should not update user when user service returns error",
			WhenRequest: &protocols.HttpRequest{
				Body: &models.UpdateUserParams{},
			},
			BeforeTest: func() error {
				userServiceMock.On("Update", mock.Anything).Return(nil, fmt.Errorf("user service threw an error"))
				return nil
			},
			ExpectResponse: func(t *testing.T, response *protocols.HttpResponse) error {
				assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
				return nil
			},
			AfterTest: func() error {
				userServiceMock.AssertCalled(t, "Update", &models.UpdateUserParams{})
				userServiceMock.AssertExpectations(t)
				userServiceMock.On("Update", mock.Anything).Unset()
				return nil
			},
		},
	}

	for _, testCase := range controllerTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			assert.NoError(t, testCase.BeforeTest())

			updateUserController := controllers.NewUpdateUserController(
				userServiceMock,
			)
			request, ok := testCase.WhenRequest.(*protocols.HttpRequest)
			assert.True(t, ok)

			response, _ := updateUserController.Handle(request)

			assert.NoError(t, testCase.ExpectResponse(t, response))
			assert.NoError(t, testCase.AfterTest())
		})
	}
}
