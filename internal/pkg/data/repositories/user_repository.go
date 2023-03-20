package repositories

import (
	"time"
	"vanir/internal/pkg/data/db"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/helpers"

	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	ID        string    `gorm:"primaryKey;type:VARCHAR(20);not null;unique"`
	Email     string    `gorm:"unique"`
	Name      string    `gorm:"type:VARCHAR(255)"`
	Password  string    `gorm:"type:VARCHAR(128)"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
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

type UserRepository struct {
	db *gorm.DB
}

var userRepository *UserRepository

func GetUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{
			db: db.GetDB(),
		}
		userRepository.db.AutoMigrate(&UserEntity{})
	}
	return userRepository
}

func (r *UserRepository) Get(ID string) (*UserEntity, error) {
	user := &UserEntity{}
	result := r.db.Where("id = ?", ID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) Create(createUserDTO *models.CreateUserDTO) (*UserEntity, error) {
	user := &UserEntity{
		Email:    createUserDTO.Email,
		Name:     createUserDTO.Name,
		Password: createUserDTO.Password,
	}
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}

func (r *UserRepository) Update(updateUserDTO *models.UpdateUserDTO) (*UserEntity, error) {

	result := r.db.Model(&UserEntity{ID: updateUserDTO.ID}).Updates(&UserEntity{
		ID:       updateUserDTO.ID,
		Name:     updateUserDTO.Name,
		Email:    updateUserDTO.Email,
		Password: updateUserDTO.NewPassword,
	})

	if result.Error != nil {
		return nil, result.Error
	} else {
		user := &UserEntity{}
		r.db.First(&user, updateUserDTO.ID)
		return user, nil
	}
}
