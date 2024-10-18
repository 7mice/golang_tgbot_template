package repository

import (
	"gorm.io/gorm"
)

type TestRepository interface {
	Test() (string, error)
}

type TestRepositoryImpl struct {
	db *gorm.DB
}

func TestRepositoryInit(db *gorm.DB) *TestRepositoryImpl {
	return &TestRepositoryImpl{
		db: db,
	}
}

func (u *TestRepositoryImpl) Test() (string, error) {

	return "test", nil
}
