package api

import (
	"IPERPUS/repository"
	"fmt"
	"net/http"
)

type controller struct {
	userRepository  repository.UserRepo
	booksRepository repository.BooksRepo
	userRegister    repository.UserRegis
	AddBooksList    repository.AddBooks
	mux             *http.ServeMux
}

func NewController(userRepository repository.UserRepo, userRegister repository.UserRegis,
	bookRepository repository.BooksRepo, AddBooksList repository.AddBooks) controller {
	mux := http.NewServeMux()
	api := controller{
		userRepository: userRepository, userRegister: userRegister, booksRepository: bookRepository, AddBooksList: AddBooksList, mux: mux,
	}
	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/signup", api.POST(http.HandlerFunc(api.signup)))
	mux.Handle("/api/bookList", api.GET(http.HandlerFunc(api.bookList))) //tambahkan tabel untuk link
	mux.Handle("/api/addBooks", api.GET(http.HandlerFunc(api.AddBooks)))

	return api

}

func (api *controller) Handler() *http.ServeMux {
	return api.mux
}

func (api *controller) Start() {
	fmt.Println("starting web server at http://localhost:2213/")
	http.ListenAndServe(":2213", api.Handler())
}
