package api

import (
	"encoding/json"
	"net/http"
	"context"
	"strings"

	"github.com/dgrijalva/jwt-go"
)
func (api *controller) AllowOrigin(w http.ResponseWriter, r *http.Request) {
	// localhost:9000 origin mendapat ijin akses
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// allow cookie
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}

func (api *controller) AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.AllowOrigin(w, r)
		encoder := json.NewEncoder(w)
		//Get token ketika request
		token := r.Header.Get("Authorization")
		if token == "" {
			// return unauthorized ketika token kosong
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: "Wrong Token"})
			return
		}
		tokenSplit := strings.Split(token, " ")
		jwtToken := tokenSplit[1]

		claims := &Claims{}

		//parse JWT token ke dalam claim
		tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// return unauthorized ketika signature invalid
				w.WriteHeader(http.StatusUnauthorized)
				encoder.Encode(AuthErrorResponse{Error: err.Error()})
				return
			}
			// return bad request ketika field token tidak ada
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}

		//return unauthorized ketika token sudah tidak valid (biasanya karna token expired)
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "role", claims.Role)
		ctx = context.WithValue(ctx, "id", claims.ID)
		ctx = context.WithValue(ctx, "props", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *controller) POST(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			encoder.Encode(AuthErrorResponse{Error: "Need POST Method!"})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *controller) GET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			encoder.Encode(AuthErrorResponse{Error: "Need GET Method!"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
