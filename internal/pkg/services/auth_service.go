package services

import (
	"fmt"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/protocols"
)

type AuthService interface {
	Authenticate(authCredentialsDTO *models.AuthCredentialsDTO) (token string, err error)
}

type AuthServiceImpl struct {
	UserRepository repositories.UserRepository
	Hasher         crypto.Hasher
	Encrypter      crypto.Encrypter
}

var authService *AuthServiceImpl

func GetAuthService() AuthService {
	if authService == nil {
		authService = &AuthServiceImpl{
			UserRepository: repositories.GetUserRepository(),
			Hasher:         crypto.GetHasher(),
			Encrypter:      crypto.GetEncrypter(),
		}
	}
	return authService
}

func (s *AuthServiceImpl) Authenticate(authCredentialsDTO *models.AuthCredentialsDTO) (string, error) {
	user, err := s.UserRepository.FindByEmail(authCredentialsDTO.Email)
	if err != nil {
		return "", err
	}

	isCompareSuccessful := s.Hasher.CompareHashes(user.Password, []byte(authCredentialsDTO.Password))

	if !isCompareSuccessful {
		return "", &protocols.AppError{
			StatusCode: 401,
			Err:        fmt.Errorf("current password is invalid"),
		}
	}

	token, err := s.Encrypter.CreateToken(user.ID)

	return token, err
}
