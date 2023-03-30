package repositories

import (
	"net/http"
	"sync"
	"time"
	"vanir/internal/pkg/data/db"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/helpers"
	"vanir/internal/pkg/protocols"

	"gorm.io/gorm"
)

type FavoriteType string

const (
	CRYPTO FavoriteType = "CRYPTO"
)

type FavoriteEntity struct {
	gorm.Model
	ID        string       `gorm:"primaryKey;type:VARCHAR(20);not null;unique"`
	Type      FavoriteType `gorm:"type:VARCHAR(32);not null"`
	Reference string       `gorm:"type:VARCHAR(255);not null"`
	UserID    string       `gorm:"type:VARCHAR(20)"`
	User      UserEntity
}

func (FavoriteEntity) TableName() string {
	return "favorites"
}

func (f *FavoriteEntity) BeforeCreate(tx *gorm.DB) error {
	f.ID = helpers.ID("u")
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	return nil
}

func (f *FavoriteEntity) BeforeUpdate(tx *gorm.DB) error {
	f.UpdatedAt = time.Now()
	return nil
}

type FavoritesRepository interface {
	Get(ID string) (*FavoriteEntity, error)
	FindAllByUserID(userID string, page int, limit int) ([]FavoriteEntity, int64, error)
	Create(params *models.CreateFavoriteParams) (*FavoriteEntity, error)
	DeleteByUserIDAndReference(userID string, reference string) error
	Delete(ID string) error
}

type FavoritesRepositoryImpl struct {
	FavoritesRepository
	db *gorm.DB
}

var favoritesRepository *FavoritesRepositoryImpl
var favoritesRepositoryOnce sync.Once

func GetFavoritesRepository() FavoritesRepository {
	favoritesRepositoryOnce.Do(func() {
		favoritesRepository = &FavoritesRepositoryImpl{
			db: db.GetDB(),
		}
		favoritesRepository.db.AutoMigrate(&FavoriteEntity{})
	})
	return favoritesRepository
}

func (r *FavoritesRepositoryImpl) Get(ID string) (*FavoriteEntity, error) {
	favorite := &FavoriteEntity{}
	result := r.db.Where("id = ?", ID).First(&favorite)
	if result.Error != nil {
		return nil, result.Error
	}
	return favorite, nil
}

func (r *FavoritesRepositoryImpl) FindAllByUserID(userID string, page int, limit int) ([]FavoriteEntity, int64, error) {
	favorites := []FavoriteEntity{}
	var count int64
	result := r.db.Where("user_id = ?", userID).Limit(limit).Offset(page * limit).Find(&favorites)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	result = r.db.Model(&FavoriteEntity{}).Where("user_id = ?", userID).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return favorites, count, nil
}

func (r *FavoritesRepositoryImpl) Create(params *models.CreateFavoriteParams) (*FavoriteEntity, error) {
	existentFavorite := &FavoriteEntity{}

	r.db.Where(&FavoriteEntity{
		UserID:    params.UserID,
		Reference: params.Reference,
	}).First(&existentFavorite)

	if existentFavorite.ID != "" {
		return existentFavorite, nil
	}

	favorite := &FavoriteEntity{
		Type:      FavoriteType(params.Type),
		Reference: params.Reference,
		UserID:    params.UserID,
	}

	result := r.db.Create(&favorite)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return favorite, nil
	}
}

func (r *FavoritesRepositoryImpl) Delete(ID string) error {
	favorite := &FavoriteEntity{ID: ID}
	result := r.db.Delete(&favorite)

	return result.Error
}

func (r *FavoritesRepositoryImpl) DeleteByUserIDAndReference(userID string, reference string) error {
	result := r.db.Where(&FavoriteEntity{UserID: userID, Reference: reference}).Delete(&FavoriteEntity{})
	if result.RowsAffected == 0 {
		return protocols.NewAppError("favorite not found", http.StatusNotFound)
	}
	return result.Error
}
