package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/domain"
	"gorm.io/gorm"
)

// RoleRepository defines the methods that any repository implementation must provide.
type RoleRepositoryInterface interface {
	Create(role *domain.Role) error
	GetByID(id uint) (*domain.Role, error)
	Update(role *domain.Role) error
	Delete(id uint) error
}

// RoleRepository handles the CRUD operations with the database.
type RoleRepository struct {
	DB *gorm.DB
}

// NewRoleRepository creates a new instance of RoleRepository.
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

// Create inserts a new role record into the database.
func (r *RoleRepository) Create(role *domain.Role) error {
	return r.DB.Create(role).Error
}

// GetByID retrieves a role record by its ID from the database.
func (r *RoleRepository) GetByID(id uint) (*domain.Role, error) {
	var role domain.Role
	err := r.DB.First(&role, id).Error
	return &role, err
}

// Update updates a role record in the database.
func (r *RoleRepository) Update(role *domain.Role) error {
	return r.DB.Save(role).Error
}

// Delete deletes a role record from the database.
func (r *RoleRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Role{}, id).Error
}
