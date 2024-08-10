package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/storage"
)

// CountryUsecaseInterface defines the methods that any use case implementation must provide.
type CountryUsecaseInterface interface {
	CreateCountry(input dto.CountryCreateDTO) error
	GetCountryByID(id uint) (*dto.CountryResponseDTO, error)
	UpdateCountry(id uint, input dto.CountryUpdateDTO) error
	DeleteCountry(id uint) error
}

// CountryUsecase handles the business logic for countries.
type CountryUsecase struct {
	CountryRepo storage.CountryRepositoryInterface
}

// NewCountryUsecase creates a new instance of CountryUsecase.
func NewCountryUsecase(CountryRepo storage.CountryRepositoryInterface) *CountryUsecase {
	return &CountryUsecase{CountryRepo: CountryRepo}
}

// CreateCountry creates a new country using the provided DTO.
func (u *CountryUsecase) CreateCountry(input dto.CountryCreateDTO) error {
	country := &domain.Country{
		Name:   input.Name,
		Status: input.Status,
	}
	err := u.CountryRepo.Create(country)
	return err
}

// GetCountryByID retrieves a country by its ID.
func (u *CountryUsecase) GetCountryByID(id uint) (*dto.CountryResponseDTO, error) {
	country, err := u.CountryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	response := &dto.CountryResponseDTO{
		Name:   country.Name,
		Status: country.Status,
	}
	return response, nil
}

// UpdateCountry updates a country using the provided DTO.
func (u *CountryUsecase) UpdateCountry(id uint, input dto.CountryUpdateDTO) error {
	country, err := u.CountryRepo.GetByID(id)
	if err != nil {
		return err
	}
	country.Name = input.Name
	country.Status = input.Status
	return u.CountryRepo.Update(country)
}

// DeleteCountry deletes a country by its ID.
func (u *CountryUsecase) DeleteCountry(id uint) error {
	return u.CountryRepo.Delete(id)
}
