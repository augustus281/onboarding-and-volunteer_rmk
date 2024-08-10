package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestRoleRepository_Create(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewRoleRepository(gormDB)

	role := &domain.Role{
		Name:   "Admin",
		Status: 123,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `roles`").WithArgs(role.Name, role.Status).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Create(role)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRoleRepository_GetByID(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewRoleRepository(gormDB)

	role := &domain.Role{
		Name:   "Admin",
		Status: 456,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).AddRow(1, role.Name, role.Status)
	mock.ExpectQuery("SELECT * FROM `roles` WHERE `roles`.`id` = ?").WithArgs(1).WillReturnRows(rows)

	result, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, role.Name, result.Name)
	assert.Equal(t, role.Status, result.Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRoleRepository_Update(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewRoleRepository(gormDB)

	role := &domain.Role{
		Name:   "Admin",
		Status: 789,
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `roles` SET `name`=?,`status`=? WHERE `id` = ?").WithArgs(role.Name, role.Status, role.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Update(role)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRoleRepository_Delete(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()
	repo := NewRoleRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `roles` WHERE `id` = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Delete(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
