package api

import (
	//	"IPERPUS/repository"
	"IPERPUS/repository"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	//	"github.com/prometheus/client_golang/api"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ResponsesTrue struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
type RegisterResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type AuthErrorResponse struct {
	Error string `json:"error"`
}

var jwtKey = []byte("key")

type Claims struct {
	Username string
	Role     string
	ID       int64
	jwt.StandardClaims
}
type Bookss struct {
	//	ID        int    `json:"id"`
	Title     string `json:"title"`
	Writer    string `json:"writer"`
	Publisher string `json:"publisher"`
}

type ListBooksSuccessResponse struct {
	book []Bookss `json :"books"`
}

func (api *controller) login(w http.ResponseWriter, req *http.Request) {

	var user User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := api.userRepository.Login(user.Username, user.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	claims := &Claims{
		Username:       *res,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
		Path:  "/",
	})
	json.NewEncoder(w).Encode(ResponsesTrue{Username: *res, Token: tokenString})

}

func (api *controller) signup(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.userRegister.Register(user.Username, user.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(*res)

}

type BookListSuccess struct {
	AllBooks []repository.Book `json: "books"`
}

func (api *controller) bookList(w http.ResponseWriter, r *http.Request) {
	//encoder := json.NewEncoder(w)
	// response := ListBooksSuccessResponse{}
	// response.book = make([]Bookss, 0)

	buku, err := api.booksRepository.ReadList()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(r) //sampai sini

	defer func() {
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
	}()

	encoder.Encode(BookListSuccess{AllBooks: buku})

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, listBuku := range buku {
	// 	response.book = append(response.book, Bookss{
	// 		//ID:        int(listBuku.ID),
	// 		Title:     listBuku.BookTitle,
	// 		Writer:    listBuku.Writer,
	// 		Publisher: listBuku.Publisher,
	// 	})
	// 	encoder.Encode(response)
	// }
	return
}

func (api *controller) AddBooks(w http.ResponseWriter, r *http.Request) {
	var book Bookss

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.AddBooksList.AddBooksList(book.Title, book.Writer, book.Publisher)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(*res)

}
