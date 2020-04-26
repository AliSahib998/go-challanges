package middleware

import (
	"context"
	"fmt"
	"github.com/AliSahib998/go-challanges/model"
	"github.com/AliSahib998/go-challanges/util"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authIgnore := []string{"/user", "/user/login", "/health"}
		requestPath := r.URL.Path
		var customError util.RestErrorResponse

		for _, value := range authIgnore {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			customError = util.RestErrorResponse{Code: 403, Message: "Missing Token"}
			util.ErrorHandler(w, customError)
			return
		}

		tokenSplit := strings.Split(tokenHeader, " ")
		if len(tokenSplit) != 2 {
			customError = util.RestErrorResponse{Code: 403, Message: "Wrong Token Header"}
			util.ErrorHandler(w, customError)
			return
		}

		if tokenSplit[0] != "Bearer" {
			customError = util.RestErrorResponse{Code: 403, Message: "Wrong Token Header"}
			util.ErrorHandler(w, customError)
			return
		}

		jwtToken := tokenSplit[1]
		tk := &model.Token{}

		token, err := jwt.ParseWithClaims(jwtToken, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("Token_Key")), nil
		})

		if err != nil {
			customError = util.RestErrorResponse{Code: 403, Message: "Malformed Authentication token"}
			util.ErrorHandler(w, customError)
			return
		}

		if !token.Valid {
			customError = util.RestErrorResponse{Code: 403, Message: "Token is wrong"}
			util.ErrorHandler(w, customError)
			return
		}

		fmt.Sprintf("username %", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.Username)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
