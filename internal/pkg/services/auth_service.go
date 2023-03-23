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
	Authenticate(authCredentialsDTO *models.AuthCredentialsDTO) (token *models.AuthenticationResponseDTO, err error)
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

func (s *AuthServiceImpl) Authenticate(authCredentialsDTO *models.AuthCredentialsDTO) (*models.AuthenticationResponseDTO, error) {
	user, err := s.userRepository.FindByEmail(authCredentialsDTO.Email)
	if err != nil {
		return nil, err
	}

	isCompareSuccessful := s.hasher.CompareHashes(user.Password, []byte(authCredentialsDTO.Password))

	if !isCompareSuccessful {
		return nil, &protocols.AppError{
			StatusCode: 401,
			Err:        fmt.Errorf("current password is invalid"),
		}
	}

	token, err := s.encrypter.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthenticationResponseDTO{
		Token: token,
	}, nil
}
