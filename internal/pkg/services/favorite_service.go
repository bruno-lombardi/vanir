package services

import (
	"math"
	"sync"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
)

type FavoriteService interface {
	Paginate(params *models.ListUserFavoritesQueryParams) (*models.ListFavoritesResponse, error)
	Create(params *models.CreateFavoriteParams) (*models.Favorite, error)
	Get(ID string) (*models.Favorite, error)
	Delete(params *models.DeleteFavoriteParams) error
}

type FavoriteServiceImpl struct {
	favoritesRepository repositories.FavoritesRepository
}

var favoriteService *FavoriteServiceImpl
var favoriteServiceOnce sync.Once

func GetFavoriteService(favoritesRepository repositories.FavoritesRepository) FavoriteService {
	favoriteServiceOnce.Do(func() {
		favoriteService = NewFavoriteServiceImpl(favoritesRepository)
	})
	return favoriteService
}

func NewFavoriteServiceImpl(favoritesRepository repositories.FavoritesRepository) *FavoriteServiceImpl {
	return &FavoriteServiceImpl{
		favoritesRepository: favoritesRepository,
	}
}

func (s *FavoriteServiceImpl) Create(params *models.CreateFavoriteParams) (*models.Favorite, error) {
	favorite, err := s.favoritesRepository.Create(params)

	if err != nil {
		return nil, err
	} else {
		return &models.Favorite{
			ID:        favorite.ID,
			Type:      string(favorite.Type),
			Reference: favorite.Reference,
			UserID:    favorite.UserID,
			CreatedAt: favorite.CreatedAt,
			UpdatedAt: favorite.UpdatedAt,
		}, nil
	}
}

func (s *FavoriteServiceImpl) Get(ID string) (*models.Favorite, error) {
	favorite, err := s.favoritesRepository.Get(ID)

	return &models.Favorite{
		ID:        favorite.ID,
		Type:      string(favorite.Type),
		Reference: favorite.Reference,
		UserID:    favorite.UserID,
		CreatedAt: favorite.CreatedAt,
		UpdatedAt: favorite.UpdatedAt,
	}, err
}

func (s *FavoriteServiceImpl) Delete(params *models.DeleteFavoriteParams) error {
	err := s.favoritesRepository.DeleteByUserIDAndReference(params.UserID, params.Reference)
	return err
}

func (s *FavoriteServiceImpl) Paginate(params *models.ListUserFavoritesQueryParams) (*models.ListFavoritesResponse, error) {
	favorites, count, err := s.favoritesRepository.FindAllByUserID(params.UserID, params.Page, params.Limit)

	return &models.ListFavoritesResponse{
		Data: mapListToModel(favorites),
		Paginated: models.Paginated{
			TotalPages: int(math.Ceil(float64(count) / float64(params.Limit))),
			Count:      int(count),
			Page:       params.Page,
			Limit:      params.Limit,
		},
	}, err
}

func mapEntityToModel(favorite repositories.FavoriteEntity) models.Favorite {
	return models.Favorite{
		ID:        favorite.ID,
		Type:      string(favorite.Type),
		Reference: favorite.Reference,
		UserID:    favorite.UserID,
		CreatedAt: favorite.CreatedAt,
		UpdatedAt: favorite.UpdatedAt,
	}
}

func mapListToModel(favorites []repositories.FavoriteEntity) []models.Favorite {
	mappedFavorites := []models.Favorite{}
	for _, item := range favorites {
		mappedFavorites = append(mappedFavorites, mapEntityToModel(item))
	}
	return mappedFavorites
}
