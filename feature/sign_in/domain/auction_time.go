package domain

import (
    "time"
    "gorm.io/gorm"
)

// AuctionTime represents the auction time data model that interacts with GORM and JSON.
type AuctionTime struct {
    ID        uint           `gorm:"primaryKey" json:"id"`        // Primary key
    StartTime time.Time      `gorm:"not null" json:"start_time"`  // Auction start time
    EndTime   time.Time      `gorm:"not null" json:"end_time"`    // Auction end time
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"` // Timestamp of creation
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"` // Timestamp of the last update
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`             // Soft delete timestamp (ignored in JSON)
}

// TableName overrides the default table name used by GORM.
func (AuctionTime) TableName() string {
    return "auction_auction_time"  // Custom table name
}
