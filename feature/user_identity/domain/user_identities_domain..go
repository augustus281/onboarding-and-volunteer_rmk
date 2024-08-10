package domain

import "time"

type UserIdentity struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int       `gorm:"not null"`
	Number      string    `gorm:"not null"`
	Type        string    `gorm:"not null"`
	Status      int       `gorm:"not null"`
	ExpiryDate  time.Time `gorm:"not null"`
	PlaceIssued string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
