package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

type UserRegis struct {
	db *sql.DB
}

func NewUserRegis(db *sql.DB) *UserRegis {
	return &UserRegis{db: db}
}

func (u *UserRepo) Login(username string, password string) (*string, error) {
	var user User
	err := u.db.QueryRow("SELECT username FROM users WHERE username = ? AND password = ?", username, password).Scan(&user.Username)
	if err != nil {
		return nil, errors.New("Login Failed")
	}
	return &user.Username, nil
}

func (u *UserRegis) Register(username string, password string) (*string, error) {
	var user User
	_, err := u.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		fmt.Println(err, "error")
		return nil, errors.New("register failed")
	}
	return &user.Username, nil
}
