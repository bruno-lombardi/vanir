package controllers_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"vanir/internal/app/presentation/controllers"
	"vanir/internal/pkg/config"
	"vanir/internal/pkg/data/db"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/helpers"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
	_ "vanir/test"
	data_test "vanir/test/data"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
				assert.Equal(t, 201, response.StatusCode)
				assert.Equal(t, user.Email, "bruno@email.com")
				return nil
			},
		},
	}

	for _, testCase := range controllerTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			assert.NoError(t, testCase.BeforeTest())

			createUserController := controllers.NewCreateUserController(
				services.GetUserService(),
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

type UserServiceMock struct {
	mock.Mock
}

func (s *UserServiceMock) Create(createUserDTO *models.CreateUserDTO) (*models.User, error) {
	args := s.Called(createUserDTO)
	user, ok := args.Get(0).(*models.User)

	if ok {
		return user, nil
	}

	err, ok := args.Get(1).(error)

	if ok {
		return nil, err
	}

	return nil, nil
}

func (s *UserServiceMock) Update(updateUserDTO *models.UpdateUserDTO) (*models.User, error) {
	args := s.Called(updateUserDTO)
	user, ok := args.Get(0).(*models.User)

	if ok {
		return user, nil
	}

	err, ok := args.Get(1).(error)

	if ok {
		return nil, err
	}

	return nil, nil
}

func TestCreateUserControllerWhenUserServiceReturnsUser(t *testing.T) {
	userServiceMock := &UserServiceMock{}
	createUserController := controllers.NewCreateUserController(
		userServiceMock,
	)
	createUserDTO := &models.CreateUserDTO{}
	request := &protocols.HttpRequest{
		Body: createUserDTO,
	}
	user := &models.User{ID: helpers.ID("u"), Email: "bruno@email.com", Password: "123123012031231", Name: "Bruno Lombardi"}

	userServiceMock.On("Create", mock.Anything).Return(user, nil)
	response, _ := createUserController.Handle(request)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	userServiceMock.AssertCalled(t, "Create", createUserDTO)
	userServiceMock.AssertExpectations(t)
}

func TestCreateUserControllerWhenUserServiceReturnsError(t *testing.T) {
	userServiceMock := &UserServiceMock{}
	createUserController := controllers.NewCreateUserController(
		userServiceMock,
	)
	createUserDTO := &models.CreateUserDTO{}
	request := &protocols.HttpRequest{
		Body: createUserDTO,
	}

	userServiceMock.On("Create", mock.Anything).Return(nil, fmt.Errorf("user service threw an error"))
	response, _ := createUserController.Handle(request)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	userServiceMock.AssertCalled(t, "Create", createUserDTO)
	userServiceMock.AssertExpectations(t)
}
