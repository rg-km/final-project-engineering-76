package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`

	CREATE TABLE books (
		title varchar(255) not null primary key,
		writer varchar(255) not null,
		publisher varchar(255) not null
	);
	
	CREATE TABLE users (
		id integer not null primary key AUTOINCREMENT,
   		username varchar(255) not null,
    	password varchar(255) not null
	);

	INSERT INTO books(title, writer, publisher) VALUES
    ('Harry Potter', 'J.K. Rowling', 'IDK');

	INSERT INTO users(username, password) VALUES
    ('vina', '1313'),
    ('kevin', '1212')`)

	if err != nil {
		panic(err)
	}
}
