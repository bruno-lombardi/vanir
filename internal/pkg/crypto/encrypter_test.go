package crypto

import (
	"fmt"
	"testing"
	"vanir/internal/pkg/config"
	"vanir/test/mocks"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JWTEncrypterSuite struct {
	suite.Suite
	encrypter *JWTEncrypter
}

func (sut *JWTEncrypterSuite) BeforeTest(_, _ string) {
	config.Setup()
	sut.encrypter = GetEncrypter().(*JWTEncrypter)
}

func (sut *JWTEncrypterSuite) TestSmokeGetEncrypter() {
	encrypter, ok := GetEncrypter().(*JWTEncrypter)
	sut.NotNil(encrypter)
	sut.NotNil(newSigner)
	sut.True(ok)
}

func (sut *JWTEncrypterSuite) TestShouldCreateTokenForSubject() {
	subject := "u_uu12312012"
	token, err := sut.encrypter.CreateToken(subject)
	sut.NotNil(token)
	sut.Nil(err)
}

func (sut *JWTEncrypterSuite) TestShouldReturnEmptyWhenSigningFails() {
	initialNewSigner := newSigner
	signerMock := &mocks.StringSignerMock{}
	mockCall := signerMock.On("SignedString", mock.Anything).Return("", fmt.Errorf("error signing string"))
	newSigner = func(method jwt.SigningMethod, claims jwt.Claims) StringSigner {
		return signerMock
	}
	defer func() {
		newSigner = initialNewSigner
		mockCall.Unset()
	}()

	subject := "u_uu12312012"
	token, err := sut.encrypter.CreateToken(subject)
	sut.NotNil(err)
	sut.Empty(token)
	signerMock.AssertExpectations(sut.T())
}

func (sut *JWTEncrypterSuite) TestShouldParseAndValidateTokenAndReturnSubject() {
	subject := "u_uu12312012"
	token, _ := sut.encrypter.CreateToken(subject)
	valid, parsed := sut.encrypter.ValidateToken(token)
	sut.Equal(subject, parsed)
	sut.True(valid)
}

func (sut *JWTEncrypterSuite) TestShouldReturnEmptySubjectWhenInvalidToken() {
	initialParseToken := parseToken
	parseToken = func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
		return nil, fmt.Errorf("error parsing token")
	}
	defer func() {
		parseToken = initialParseToken
	}()

	token := "invalid_token"
	valid, parsed := sut.encrypter.ValidateToken(token)
	sut.Empty(parsed)
	sut.False(valid)
}

func TestJWTEncrypterSuite(t *testing.T) {
	suite.Run(t, &JWTEncrypterSuite{})
}
