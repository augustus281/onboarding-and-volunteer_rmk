package dto

type ApplicantCreateDTO struct {
	Email   string `json:"email" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}

type ApplicantUpdateDTO struct {
	Email             string `json:"email"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Gender            string `json:"gender"`
	DOB               string `json:"dob"`
	Mobile            string `json:"mobile"`
	RoleID            int    `json:"role_id"`
	CountryID         int    `json:"country_id"`
	ResidentCountryID int    `json:"resident_country_id"`
	DepartmentID      int    `json:"department_id"`
}

type ApplicantResponseDTO struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Gender            string `json:"gender"`
	DOB               string `json:"dob"`
	Mobile            string `json:"mobile"`
	RoleID            int    `json:"role_id"`
	CountryID         int    `json:"country_id"`
	ResidentCountryID int    `json:"resident_country_id"`
	DepartmentID      int    `json:"department_id"`
}
