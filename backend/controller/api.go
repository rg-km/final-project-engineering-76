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
	addBooksList    repository.AddBooks
	findingBook     repository.Finder
	profilUserRepo  repository.ProfilUserReposistory
	profiladminRepo repository.ProfiladminReposistory

	mux *http.ServeMux
}

func NewController(userRepository repository.UserRepo, booksRepository repository.BooksRepo,
	userRegister repository.UserRegis, addBooksList repository.AddBooks, findingBook repository.Finder, 
	profilUserRepo  repository.ProfilUserReposistory, profiladminRepo repository.ProfiladminReposistory) controller {
	mux := http.NewServeMux()
	api := controller{
		userRepository: userRepository, booksRepository: booksRepository, userRegister: userRegister,
		addBooksList: addBooksList, findingBook: findingBook, profilUserRepo:profilUserRepo, profiladminRepo: profiladminRepo, mux: mux,
	}
	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/signup", api.POST(http.HandlerFunc(api.signup)))
	mux.Handle("/api/bookList", api.GET(http.HandlerFunc(api.bookList))) //tambahkan tabel untuk link
	mux.Handle("/api/addBooks", api.GET(http.HandlerFunc(api.addBooks)))
	mux.Handle("/api/findBooksWriter", api.GET(http.HandlerFunc(api.findBooksWriter)))
	mux.Handle("/api/findBooksTahun", api.GET(http.HandlerFunc(api.findBooksTahun)))
	mux.Handle("/api/findBooksKategori", api.GET(http.HandlerFunc(api.findBooksKategori)))
	mux.Handle("/api/dashboardUser", api.GET(http.HandlerFunc(api.dashboardUser)))
	mux.Handle("/api/dashboardAdmin", api.GET(http.HandlerFunc(api.dashboadAdmin)))
	mux.Handle("/api/profiluser", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.addProfilUser))))
	mux.Handle("/api/profiluser/edit", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.editProfilUser))))
	mux.Handle("/api/profiladmin", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.addProfilAdmin))))
	mux.Handle("/api/profiladmin/edit", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.editProfiladmin))))

	return api

}

func (api *controller) Handler() *http.ServeMux {
	return api.mux
}

func (api *controller) Start() {
	fmt.Println("starting web server at http://localhost:2213/")
	http.ListenAndServe(":2213", api.Handler())
}
