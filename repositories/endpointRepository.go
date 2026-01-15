package repositories

import (
	"github.com/SaraLuciaa/ssl-labs-challenge/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EndpointRepository interface {
	Create(endpoint *models.Endpoint) error
	Update(endpoint *models.Endpoint) error
	FindByAnalysisIDAndIP(analysisID uuid.UUID, ipAddress string) (*models.Endpoint, error)
	DeleteByAnalysisID(analysisID uuid.UUID) error
}

type endpointRepository struct {
	db *gorm.DB
}

func NewEndpointRepository(db *gorm.DB) EndpointRepository {
	return &endpointRepository{
		db: db,
	}
}

func (r *endpointRepository) Create(endpoint *models.Endpoint) error {
	return r.db.Create(endpoint).Error
}

func (r *endpointRepository) Update(endpoint *models.Endpoint) error {
	return r.db.Save(endpoint).Error
}

func (r *endpointRepository) FindByAnalysisIDAndIP(analysisID uuid.UUID, ipAddress string) (*models.Endpoint, error) {
	var endpoint models.Endpoint
	err := r.db.Where("analysis_id = ? AND ip_address = ?", analysisID, ipAddress).First(&endpoint).Error
	if err != nil {
		return nil, err
	}
	return &endpoint, nil
}

func (r *endpointRepository) DeleteByAnalysisID(analysisID uuid.UUID) error {
	return r.db.Where("analysis_id = ?", analysisID).Delete(&models.Endpoint{}).Error
}
