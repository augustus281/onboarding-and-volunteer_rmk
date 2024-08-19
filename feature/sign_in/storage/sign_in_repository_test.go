package storage

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/domain"
)

func setupTestDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    if err := db.AutoMigrate(&domain.SignIn{}); err != nil {
        return nil, err
    }
    return db, nil
}

func TestCreateSignIn(t *testing.T) {
    db, err := setupTestDB()
    if err != nil {
        t.Fatalf("could not set up test DB: %v", err)
    }
    repo := NewSignInRepository(db)

    signIn := &domain.SignIn{Username: "testuser", Email: "test@example.com"}

    err = repo.CreateSignIn(signIn)
    assert.NoError(t, err)
}

func TestGetSignInByUsername(t *testing.T) {
    db, err := setupTestDB()
    if err != nil {
        t.Fatalf("could not set up test DB: %v", err)
    }
    repo := NewSignInRepository(db)

    signIn := &domain.SignIn{Username: "testuser", Email: "test@example.com"}
    repo.CreateSignIn(signIn)

    result, err := repo.GetSignInByUsername("testuser")
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "testuser", result.Username)
}

func TestGetSignInByEmail(t *testing.T) {
    db, err := setupTestDB()
    if err != nil {
        t.Fatalf("could not set up test DB: %v", err)
    }
    repo := NewSignInRepository(db)

    signIn := &domain.SignIn{Username: "testuser", Email: "test@example.com"}
    repo.CreateSignIn(signIn)

    result, err := repo.GetSignInByEmail("test@example.com")
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "test@example.com", result.Email)
}

func TestUpdateSignIn(t *testing.T) {
    db, err := setupTestDB()
    if err != nil {
        t.Fatalf("could not set up test DB: %v", err)
    }
    repo := NewSignInRepository(db)

    signIn := &domain.SignIn{Username: "testuser", Email: "test@example.com"}
    repo.CreateSignIn(signIn)

    signIn.Email = "new-email@example.com"
    err = repo.UpdateSignIn(signIn)
    assert.NoError(t, err)

    updatedSignIn, err := repo.GetSignInByEmail("new-email@example.com")
    assert.NoError(t, err)
    assert.NotNil(t, updatedSignIn)
    assert.Equal(t, "new-email@example.com", updatedSignIn.Email)
}

func TestDeleteSignIn(t *testing.T) {
    db, err := setupTestDB()
    if err != nil {
        t.Fatalf("could not set up test DB: %v", err)
    }
    repo := NewSignInRepository(db)

    signIn := &domain.SignIn{Username: "testuser", Email: "test@example.com"}
    repo.CreateSignIn(signIn)

    err = repo.DeleteSignIn(signIn.ID)
    assert.NoError(t, err)

    result, err := repo.GetSignInByEmail("test@example.com")
    assert.NoError(t, err)
    assert.Nil(t, result)
}
