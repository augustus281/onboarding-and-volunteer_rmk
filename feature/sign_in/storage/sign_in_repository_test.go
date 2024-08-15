package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateSignIn(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err)

	repo := NewSignInRepository(gormDB)

	signIn := &domain.SignIn{
		Username:     "testuser",
		PasswordHash: "hashedpassword",
		Email:        "test@example.com",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"sign_ins\"").WithArgs(signIn.Username, signIn.PasswordHash, signIn.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.CreateSignIn(signIn)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetSignInByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err)

	repo := NewSignInRepository(gormDB)

	username := "testuser"
	rows := sqlmock.NewRows([]string{"id", "username", "password_hash", "email"}).
		AddRow(1, username, "hashedpassword", "test@example.com")

	mock.ExpectQuery("SELECT * FROM \"sign_ins\" WHERE username = ?").WithArgs(username).WillReturnRows(rows)

	signIn, err := repo.GetSignInByUsername(username)
	require.NoError(t, err)
	assert.NotNil(t, signIn)
	assert.Equal(t, username, signIn.Username)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetSignInByUsername_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err)

	repo := NewSignInRepository(gormDB)

	username := "nonexistentuser"
	mock.ExpectQuery("SELECT * FROM \"sign_ins\" WHERE username = ?").WithArgs(username).WillReturnRows(sqlmock.NewRows([]string{}))

	signIn, err := repo.GetSignInByUsername(username)
	assert.NoError(t, err)
	assert.Nil(t, signIn)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetSignInByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err)

	repo := NewSignInRepository(gormDB)

	email := "test@example.com"
	rows := sqlmock.NewRows([]string{"id", "username", "password_hash", "email"}).
		AddRow(1, "testuser", "hashedpassword", email)

	mock.ExpectQuery("SELECT * FROM \"sign_ins\" WHERE email = ?").WithArgs(email).WillReturnRows(rows)

	signIn, err := repo.GetSignInByEmail(email)
	require.NoError(t, err)
	assert.NotNil(t, signIn)
	assert.Equal(t, email, signIn.Email)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateSignIn(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err)

	repo := NewSignInRepository(gormDB)

	signIn := &domain.SignIn{
		ID:           1,
		Username:     "testuser",
		PasswordHash: "newhashedpassword",
		Email:        "test@example.com",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"sign_ins\"").WithArgs(signIn.Username, signIn.PasswordHash, signIn.Email, signIn.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.UpdateSignIn(signIn)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
