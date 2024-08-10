package dto

// CountryCreateDTO represents the data transfer object for creating a country.
type CountryCreateDTO struct {
	Name   string `json:"name" binding:"required"`
	Status uint   `json:"status" binding:"required"`
}

// CountryUpdateDTO represents the data transfer object for updating a country.
type CountryUpdateDTO struct {
	Name   string `json:"name"`
	Status uint   `json:"status"`
}

// CountryResponseDTO represents the data transfer object for a country response.
type CountryResponseDTO struct {
	Name   string `json:"name"`
	Status uint   `json:"status"`
}
