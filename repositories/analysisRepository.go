package repositories

import (
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnalysisRepository interface {
	Create(analysis *models.Analysis) error
	Update(analysis *models.Analysis) error
	FindByID(id uuid.UUID) (*models.Analysis, error)
	FindAll() ([]models.Analysis, error)
}

type analysisRepository struct {
	db *gorm.DB
}

func NewAnalysisRepository(db *gorm.DB) AnalysisRepository {
	return &analysisRepository{
		db: db,
	}
}

func (r *analysisRepository) Create(analysis *models.Analysis) error {
	return r.db.Create(analysis).Error
}

func (r *analysisRepository) Update(analysis *models.Analysis) error {
	return r.db.Save(analysis).Error
}

func (r *analysisRepository) FindByID(id uuid.UUID) (*models.Analysis, error) {
	var analysis models.Analysis
	err := r.db.Preload("Endpoints").First(&analysis, "id = ?", id).Error
	return &analysis, err
}

func (r *analysisRepository) FindAll() ([]models.Analysis, error) {
	var analyses []models.Analysis
	err := r.db.Preload("Endpoints").Order("created_at DESC").Find(&analyses).Error
	return analyses, err
}
