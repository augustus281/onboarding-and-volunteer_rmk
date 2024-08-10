package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.ApplicantDomain{}, &domain.Request{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
