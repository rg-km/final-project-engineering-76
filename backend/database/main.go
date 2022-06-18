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

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name varchar(255) NOT NULL,
		email varchar(255) NOT NULL,
		password varchar(255) NOT NULL
	);


	CREATE TABLE books (
		id integer not null primary key AUTOINCREMENT,
		book_title varchar(255) not null,
		writer varchar(255) not null,
		publisher varchar(255) not null
	);
	
	CREATE TABLE users (
    	id integer not null primary key AUTOINCREMENT,
   		username varchar(255) not null,
    	password varchar(255) not null,
    	loggedin boolean not null
	);

	INSERT INTO books(book_title, writer, publisher) VALUES
    ('Harry Potter', 'J.K. Rowling', 'IDK');

	INSERT INTO users(username, password, loggedin) VALUES
    ('vina', '1313', false),
    ('kevin', '1212', false)
	
	INSERT INTO users(username, password) VALUES
	('reza', 'real2123', false),
	('amel', 'bisayuk', false),`)



	if err != nil {
		panic(err)
	}
}
