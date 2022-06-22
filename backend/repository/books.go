package repository

import (
	"database/sql"
	"errors"
	"fmt"
	//"errors"
)

type BooksRepo struct {
	db *sql.DB
}

//type buku struct{Book}
func NewBooksRepository(db *sql.DB) *BooksRepo {
	return &BooksRepo{db: db}
}

type AddBooks struct {
	db *sql.DB
}

func AddNewBooks(db *sql.DB) *AddBooks {
	return &AddBooks{db: db}
}

func (u *BooksRepo) ReadList() ([]Book, error) {
	rows, err := u.db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var sbuku []Book

	for rows.Next() {
		var pbuku Book
		err := rows.Scan(&pbuku.BookTitle, &pbuku.Writer, &pbuku.Publisher)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sbuku = append(sbuku, pbuku)
	}
	return sbuku, nil
}

func (u *AddBooks) AddBooksList(title string, writer string, publisher string) (*string, error) {
	var buku Book
	_, err := u.db.Exec("INSERT INTO books (title,writer,publisher) VALUES(?,?,?)", title, writer, publisher)
	if err != nil {
		fmt.Println(err, "error")
		return nil, errors.New("Failed to add book")
	}
	return &buku.BookTitle, nil
}
