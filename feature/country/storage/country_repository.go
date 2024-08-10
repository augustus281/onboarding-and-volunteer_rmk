package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/domain"
	"gorm.io/gorm"
)

// CountryRepositoryInterface defines the methods that any repository implementation must provide.
type CountryRepositoryInterface interface {
	Create(country *domain.Country) error
	GetByID(id uint) (*domain.Country, error)
	Update(country *domain.Country) error
	Delete(id uint) error
}

// CountryRepository handles the CRUD operations with the database.
type CountryRepository struct {
	DB *gorm.DB
}

// NewCountryRepository creates a new instance of CountryRepository.
func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{DB: db}
}

// Create inserts a new country record into the database.
func (r *CountryRepository) Create(country *domain.Country) error {
	return r.DB.Create(country).Error
}

// GetByID retrieves a country record by its ID from the database.
func (r *CountryRepository) GetByID(id uint) (*domain.Country, error) {
	var country domain.Country
	err := r.DB.First(&country, id).Error
	return &country, err
}

// Update updates a country record in the database.
func (r *CountryRepository) Update(country *domain.Country) error {
	return r.DB.Save(country).Error
}

// Delete deletes a country record from the database.
func (r *CountryRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Country{}, id).Error
}
