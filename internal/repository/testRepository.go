package repository

import (
	"github.com/jmoiron/sqlx"
)

type TestRepository interface {
	Test() (string, error)
}

type TestRepositoryImpl struct {
	db *sqlx.DB
}

func TestRepositoryInit(db *sqlx.DB) *TestRepositoryImpl {
	return &TestRepositoryImpl{
		db: db,
	}
}

func (u *TestRepositoryImpl) Test() (string, error) {

	return "test", nil
}
