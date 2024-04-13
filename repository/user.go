package repository

import (
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/logger"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository/models"
	"github.com/google/uuid"
)

// UserRepo is an interface that defines the functions that a User repository should implement.
// It should not have any dependencies on the gin package.
type UserRepo interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetOne(id uuid.UUID) (models.User, error)
	Update(id uuid.UUID, user *models.User) error
}

func NewUserRepo(gorm *database.GormDatabase) UserRepo {
	return &userRepo{gorm: gorm}
}

type userRepo struct {
	gorm *database.GormDatabase
}

func (u *userRepo) Create(userModel *models.User) error {
	if err := u.gorm.DB.Create(&userModel).Error; err != nil {
		logger.Errorf("error: %v", err)
		return err
	}
	return nil
}

func (u *userRepo) GetAll() ([]models.User, error) {
	var users []models.User
	if err := u.gorm.DB.Raw(
		"SELECT * FROM users",
	).Scan(&users).Error; err != nil {
		logger.Errorf("error: %v", err)
	}

	return users, nil
}

func (u *userRepo) GetOne(id uuid.UUID) (models.User, error) {
	var user models.User
	if err := u.gorm.DB.First(&user, id).Error; err != nil {
		logger.Errorf("error: %v", err)
		return user, err
	}
	return user, nil
}

func (u *userRepo) Update(id uuid.UUID, user *models.User) error {
	if err := u.gorm.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		logger.Errorf("error: %v", err)
		return err
	}
	return nil
}
