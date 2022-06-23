package main

import (
	api "IPERPUS/controller"
	"IPERPUS/repository"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	usersRepo := repository.NewUserRepo(db)
	regisRepo := repository.NewUserRegis(db)
	booksRepo := repository.NewBooksRepository(db)
	addBooks := repository.AddNewBooks(db)
	profilUserRepo := repository.NewProfilUserRepository(db)
	profiladminRepo := repository.NewProfiladminRepository(db)

	mainApi := api.NewController(*usersRepo, *regisRepo, *booksRepo, *addBooks, *profilUserRepo, *profiladminRepo)
	mainApi.Start()
}
