package storage

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to open gorm db: %v", err)
	}

	cleanup := func() {
		db.Close()
	}

	return gdb, mock, cleanup
}

func TestGetListPendingRequest(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAdminRepository(db)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "requests" WHERE status = $1`)).
		WithArgs(0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "type", "status"}).
			AddRow(1, 1, "registration", 0))

	requests, err := repo.GetListPendingRequest()
	if err != "" {
		t.Errorf("unexpected error: %v", err)
	}

	if len(requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(requests))
	}
}

func TestGetPendingRequestByID(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAdminRepository(db)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "requests" WHERE id = $1 AND status = $2 ORDER BY "requests"."id" LIMIT 1`)).
		WithArgs(1, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "type", "status"}).
			AddRow(1, 1, "registration", 0))

	request, err := repo.GetPendingRequestByID(1)
	if err != "" {
		t.Errorf("unexpected error: %v", err)
	}

	if request.ID != 1 {
		t.Errorf("expected request ID to be 1, got %d", request.ID)
	}
}

func TestGetListAllRequest(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewAdminRepository(&mockDB.DB)

	expectedRequests := []*domain.Request{
		{ID: 1, UserID: 1, Type: "type1", Status: 0},
		{ID: 2, UserID: 2, Type: "type2", Status: 1},
	}
	mockDB.On("Find", &[]*domain.Request{}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]*domain.Request)
		*arg = expectedRequests
	}).Return(mockDB)

	result, msg := repo.GetListAllRequest()

	assert.Equal(t, expectedRequests, result)
	assert.Empty(t, msg)
	mockDB.AssertExpectations(t)
}

func TestGetRequestByID(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewAdminRepository(&mockDB.DB)

	expectedRequest := &domain.Request{ID: 1, UserID: 1, Type: "type1", Status: 0}
	mockDB.On("Where", "id = ?", 1).Return(mockDB)
	mockDB.On("First", &domain.Request{}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.Request)
		*arg = *expectedRequest
	}).Return(mockDB)

	result, msg := repo.GetRequestByID(1)

	assert.Equal(t, expectedRequest, result)
	assert.Empty(t, msg)
	mockDB.AssertExpectations(t)
}

func TestApproveRequest(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAdminRepository(db)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "requests" WHERE id = $1 ORDER BY "requests"."id" LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "type", "status"}).
			AddRow(1, 1, "registration", 0))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "requests" SET "status"=$1,"verifier_id"=$2 WHERE "id" = $3`)).
		WithArgs(1, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "role_id"=$1 WHERE "id" = $2`)).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.ApproveRequest(1, 1)
	if err != "" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRejectRequest(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAdminRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "requests" SET "status"=$1,"verifier_id"=$2 WHERE "id" = $3`)).
		WithArgs(2, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.RejectRequest(1, 1)
	if err != "" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAddRejectNotes(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAdminRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "requests" SET "reject_notes"=$1 WHERE "id" = $2`)).
		WithArgs("some notes", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.AddRejectNotes(1, "some notes")
	if err != "" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDeleteRequest(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAdminRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "requests" WHERE "id" = $1`)).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteRequest(1)
	if err != "" {
		t.Errorf("unexpected error: %v", err)
	}
}
