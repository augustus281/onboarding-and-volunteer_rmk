package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	dialector := mysql.New(mysql.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return gormDB, mock, nil
}

func TestCountryRepository_Create(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewCountryRepository(gormDB)

	country := &domain.Country{
		Name: "Test Country",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO countries`).
		WithArgs(sqlmock.AnyArg(), country.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Create(country)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountryRepository_GetByID(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewCountryRepository(gormDB)

	countryID := uint(1)
	country := &domain.Country{
		Id:   countryID,
		Name: "Test Country",
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(country.Id, country.Name)

	mock.ExpectQuery(`SELECT \* FROM countries WHERE id = \?`).
		WithArgs(countryID).
		WillReturnRows(rows)

	result, err := repo.GetByID(countryID)
	assert.NoError(t, err)
	assert.Equal(t, country, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountryRepository_Update(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewCountryRepository(gormDB)

	country := &domain.Country{
		Id:   1,
		Name: "Updated Country",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE countries SET name = \?, code = \? WHERE id = \?`).
		WithArgs(country.Name, country.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Update(country)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountryRepository_Delete(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewCountryRepository(gormDB)

	countryID := uint(1)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM countries WHERE id = \?`).
		WithArgs(countryID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Delete(countryID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
