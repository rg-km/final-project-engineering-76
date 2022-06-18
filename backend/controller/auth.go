package api

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponseTrue struct {
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
	jwt.StandardClaims 
	
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
	json.NewEncoder(w).Encode(LoginResponseTrue{Username: *res, Token: tokenString})

}
func (api *controller) register(w http.ResponseWriter, req *http.Request) {

	var user User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := api.userRepository.Register(user.Username, user.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(RegisterResponse{Username: *res, Password: user.Password})
	w.WriteHeader(http.StatusOK)
}



	
	
	
