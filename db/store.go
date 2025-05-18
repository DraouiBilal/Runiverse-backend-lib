package db

import (
	"database/sql"
)

type Store[T any] interface {
	GetAll(db *sql.DB) ([]T, error)
	GetById(db *sql.DB, id string) (T, error)
	Create(db *sql.DB, item T) (T, error)
	Update(db *sql.DB, item T) (T, error)
	Delete(db *sql.DB, item T) error
}
