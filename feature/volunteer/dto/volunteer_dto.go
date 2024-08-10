package dto

type VolunteerCreateDTO struct {
	UserID       int `json:"user_id" binding:"required"`
	DepartmentID int `json:"department_id" binding:"required"`
	Status       int `json:"status" binding:"required"`
}

type VolunteerUpdateDTO struct {
	DepartmentID int `json:"department_id"`
	Status       int `json:"status"`
}

type VolunteerResponseDTO struct {
	ID           int `json:"id"`
	UserID       int `json:"user_id"`
	DepartmentID int `json:"department_id"`
	Status       int `json:"status"`
}
