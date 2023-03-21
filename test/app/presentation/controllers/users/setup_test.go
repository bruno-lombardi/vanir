package user_controllers_test

import (
	"os"
	"testing"
	"vanir/internal/pkg/config"
	"vanir/internal/pkg/data/db"
	data_test "vanir/test/data"
	"vanir/test/data/mocks"
)

type ControllerTestCase = data_test.ControllerTestCase

var userServiceMock *mocks.UserServiceMock

func TestMain(m *testing.M) {
	config.Setup()
	db.SetupDB()
	userServiceMock = &mocks.UserServiceMock{}

	code := m.Run()

	os.Exit(code)
}
