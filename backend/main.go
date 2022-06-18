package main

import (
	api "IPERPUS/controller"
	"IPERPUS/repository"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}
	usersRepo := repository.NewUserRepo(db)
	//booksRepo := repository.NewBooksRepository(db)

	mainApi := api.NewController(*usersRepo)
	mainApi.Start()
}
