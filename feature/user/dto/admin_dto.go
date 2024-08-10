package dto

import (
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
)

type PendingRequest struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Type     string    `json:"type"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type RequestResponse struct {
	ID          int       `json:"id"`
	UserID      uint      `json:"user_id"`
	Type        string    `json:"type"`
	Status      int       `json:"status"`
	RejectNotes string    `json:"reject_notes"`
	VerifierID  int       `json:"verifier_id"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type ListRequest struct {
	Requests []*domain.Request `json:"requests"`
}

type AddRejectNoteRequest struct {
	Notes string `json:"notes"`
}
