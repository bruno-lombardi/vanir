package data_test

import (
	"testing"
	"vanir/internal/pkg/protocols"
)

type ControllerTestCase struct {
	Name           string
	WhenRequest    interface{}
	BeforeTest     func() error
	ExpectResponse func(t *testing.T, response *protocols.HttpResponse) error
	AfterTest      func() error
}
