package controllers_test

import (
	"os"
	"testing"
	"vanir/internal/app/presentation/controllers"
	"vanir/internal/pkg/config"
	"vanir/internal/pkg/data/db"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
	_ "vanir/test"
	data_test "vanir/test/data"

	"github.com/stretchr/testify/assert"
)

type ControllerTestCase = data_test.ControllerTestCase

func TestMain(m *testing.M) {
	config.Setup()
	db.SetupDB()

	code := m.Run()

	os.Exit(code)
}

func TestCreateUserController(t *testing.T) {
	controllerTestCases := []ControllerTestCase{
		{
			Name: "Should create user with valid data",
			WhenRequest: &protocols.HttpRequest{
				Body: &models.CreateUserDTO{
					Name:                 "Bruno Lombardi",
					Email:                "bruno@email.com",
					Password:             "123456",
					PasswordConfirmation: "123456",
				},
			},
			BeforeTest: func() error { return nil },
			AfterTest:  func() error { return nil },
			ExpectResponse: func(t *testing.T, response *protocols.HttpResponse) error {
				user, ok := response.Body.(*models.User)
				assert.True(t, ok)
				assert.NotNil(t, user.ID)
				assert.Equal(t, user.Email, "bruno@email.com")
				return nil
			},
		},
	}

	for _, testCase := range controllerTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			assert.NoError(t, testCase.BeforeTest())

			createUserController := controllers.NewCreateUserController(
				*services.GetUserService(),
			)
			request, ok := testCase.WhenRequest.(*protocols.HttpRequest)
			assert.True(t, ok)

			response, err := createUserController.Handle(request)

			assert.NoError(t, err)
			assert.NoError(t, testCase.ExpectResponse(t, response))
			assert.NoError(t, testCase.AfterTest())
		})
	}
}
