package api

import (
	"fmt"

	"IPERPUS/repository"
	"net/http"
)

type controller struct {
	userRepository  repository.UserRepo
	registerRepository repository.RegisterRepo 
	booksRepository repository.BooksRepo
	mux             *http.ServeMux
}

func NewController(userRepository repository.UserRepo) controller {
	mux := http.NewServeMux()
	api := controller{
		userRepository: userRepository, mux: mux,
	}
	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/register", api.POST(http.HandlerFunc(api.register)))

	return api

}

func (api *controller) Handler() *http.ServeMux {
	return api.mux
}

func (api *controller) Start() {
	fmt.Println("starting web server at http://localhost:2213/")
	http.ListenAndServe(":2213", api.Handler())
}
