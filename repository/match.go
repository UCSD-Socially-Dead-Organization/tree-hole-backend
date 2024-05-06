package repository

import (
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/database"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/infra/logger"
	"github.com/UCSD-Socially-Dead-Organization/tree-hole-backend/repository/models"
	"github.com/google/uuid"
)

// UserRepo is an interface that defines the functions that a User repository should implement.
// It should not have any dependencies on the gin package.
type MatchRepo interface {
	Create(match *models.Match) error
	GetAll() ([]models.Match, error)
	GetOne(id uuid.UUID) (models.Match, error)
	Update(match *models.Match) error
}

func NewMatchRepo(gorm *database.GormDatabase) MatchRepo {
	return &matchRepo{gorm: gorm}
}

type matchRepo struct {
	gorm *database.GormDatabase
}

func (r *matchRepo) Create(matchModel *models.Match) error {
	if err := r.gorm.DB.Create(&matchModel).Error; err != nil {
		logger.Errorf("error: %v", err)
		return err
	}
	return nil
}

func (r *matchRepo) GetAll() ([]models.Match, error) {
	var users []models.Match
	if err := r.gorm.DB.Raw(
		"SELECT * FROM matches",
	).Scan(&users).Error; err != nil {
		logger.Errorf("error: %v", err)
	}

	return users, nil
}

func (r *matchRepo) GetOne(id uuid.UUID) (models.Match, error) {
	var user models.Match
	if err := r.gorm.DB.First(&user, id).Error; err != nil {
		logger.Errorf("error: %v", err)
		return user, err
	}
	return user, nil
}

func (r *matchRepo) Update(match *models.Match) error {
	if err := r.gorm.DB.Model(&models.Match{}).Where("id = ?", match.ID).Updates(match).Error; err != nil {
		logger.Errorf("error: %v", err)
		return err
	}
	return nil
}
