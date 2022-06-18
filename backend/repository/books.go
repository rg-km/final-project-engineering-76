package repository

import (
	"database/sql"
	//"errors"
)

type BooksRepo struct {
	db *sql.DB
}

func NewBooksRepository(db *sql.DB) *BooksRepo {
	return &BooksRepo{db: db}
}
