package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"gorm.io/gorm"
)

// DepartmentRepository defines the methods that any repository implementation must provide.
type DepartmentRepositoryInterface interface {
	Create(department *domain.Department) error
	GetByID(id uint) (*domain.Department, error)
	Update(department *domain.Department) error
	Delete(id uint) error
}

// DepartmentRepository handles the CRUD operations with the database.
type DepartmentRepository struct {
	DB *gorm.DB
}

// NewDepartmentRepository creates a new instance of DepartmentRepository.
func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{DB: db}
}

// Create inserts a new department record into the database.
func (r *DepartmentRepository) Create(department *domain.Department) error {
	return r.DB.Create(department).Error
}

// GetByID retrieves a department record by its ID from the database.
func (r *DepartmentRepository) GetByID(id uint) (*domain.Department, error) {
	var department domain.Department
	err := r.DB.First(&department, id).Error
	return &department, err
}

// Update updates a department record in the database.
func (r *DepartmentRepository) Update(department *domain.Department) error {
	return r.DB.Save(department).Error
}

// Delete deletes a department record from the database.
func (r *DepartmentRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Department{}, id).Error
}
