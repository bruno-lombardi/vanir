package repositories

import (
	"sync"
	"time"
	"vanir/internal/pkg/data/db"
	"vanir/internal/pkg/helpers"

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
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
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
	FindByEmail(email string) (*FavoriteEntity, error)
	Create(params string) (*FavoriteEntity, error)
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
		favoritesRepository.db.AutoMigrate(&UserEntity{})
	})
	return favoritesRepository
}
