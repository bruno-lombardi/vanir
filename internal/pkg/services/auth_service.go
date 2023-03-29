package services

import (
	"fmt"
	"sync"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/protocols"
)

type AuthService interface {
	Authenticate(authCredentials *models.AuthCredentials) (token *models.AuthenticationResponse, err error)
}

type AuthServiceImpl struct {
	userRepository repositories.UserRepository
	hasher         crypto.Hasher
	encrypter      crypto.Encrypter
}

var authService *AuthServiceImpl
var authServiceOnce sync.Once

func GetAuthService() AuthService {
	authServiceOnce.Do(func() {
		authService = &AuthServiceImpl{
			userRepository: repositories.GetUserRepository(),
			hasher:         crypto.GetHasher(),
			encrypter:      crypto.GetEncrypter(),
		}
	})
	return authService
}

func NewAuthServiceImpl(userRepository repositories.UserRepository, hasher crypto.Hasher, encrypter crypto.Encrypter) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository,
		hasher:         hasher,
		encrypter:      encrypter,
	}
}

func (s *AuthServiceImpl) Authenticate(authCredentials *models.AuthCredentials) (*models.AuthenticationResponse, error) {
	user, err := s.userRepository.FindByEmail(authCredentials.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	isCompareSuccessful := s.hasher.CompareHashes(user.Password, []byte(authCredentials.Password))

	if !isCompareSuccessful {
		return nil, &protocols.AppError{
			StatusCode: 401,
			Err:        fmt.Errorf("invalid credentials"),
		}
	}

	token, err := s.encrypter.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthenticationResponse{
		Token: token,
	}, nil
}
