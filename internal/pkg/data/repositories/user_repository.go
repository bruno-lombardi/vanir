package repositories

import (
	"sync"
	"time"
	"vanir/internal/pkg/data/db"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/helpers"

	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	ID        string           `gorm:"primaryKey;type:VARCHAR(20);not null;unique"`
	Email     string           `gorm:"unique"`
	Name      string           `gorm:"type:VARCHAR(255)"`
	Password  string           `gorm:"type:VARCHAR(128)"`
	Favorites []FavoriteEntity `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time        `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
	UpdatedAt time.Time        `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
}

func (UserEntity) TableName() string {
	return "users"
}

func (u *UserEntity) BeforeCreate(tx *gorm.DB) error {
	u.ID = helpers.ID("u")
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *UserEntity) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

type UserRepository interface {
	Get(ID string) (*UserEntity, error)
	FindByEmail(email string) (*UserEntity, error)
	Create(createUserParams *models.CreateUserParams) (*UserEntity, error)
	Update(updateUserParams *models.UpdateUserParams) (*UserEntity, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

var userRepository *UserRepositoryImpl
var userRepositoryOnce sync.Once

func GetUserRepository() UserRepository {
	userRepositoryOnce.Do(func() {
		userRepository = &UserRepositoryImpl{
			db: db.GetDB(),
		}
		userRepository.db.AutoMigrate(&UserEntity{})
	})
	return userRepository
}

func (r *UserRepositoryImpl) Get(ID string) (*UserEntity, error) {
	user := &UserEntity{}
	result := r.db.Where("id = ?", ID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*UserEntity, error) {
	user := &UserEntity{}
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) Create(createUserParams *models.CreateUserParams) (*UserEntity, error) {
	user := &UserEntity{
		Email:    createUserParams.Email,
		Name:     createUserParams.Name,
		Password: createUserParams.Password,
	}
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}

func (r *UserRepositoryImpl) Update(updateUserParams *models.UpdateUserParams) (*UserEntity, error) {

	result := r.db.Model(&UserEntity{ID: updateUserParams.ID}).Updates(&UserEntity{
		ID:       updateUserParams.ID,
		Name:     updateUserParams.Name,
		Email:    updateUserParams.Email,
		Password: updateUserParams.NewPassword,
	})

	if result.Error != nil {
		return nil, result.Error
	} else {
		user, err := r.Get(updateUserParams.ID)
		return user, err
	}
}
