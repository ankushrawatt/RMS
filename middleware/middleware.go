package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"rmsProject/model"
	"rmsProject/utils"
)

type ContextKeys string

const (
	userContext ContextKeys = "__userContext"
)

var mySigningKey = []byte("secret_key")

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apikey := request.Header.Get("x-api-key")
		claims := model.JWTClaims{}
		token, tokenErr := jwt.ParseWithClaims(apikey, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			//fmt.Println(mySigningKey)
			return mySigningKey, nil
		})
		if tokenErr != nil {
			utils.CheckError(tokenErr)
		}
		if token.Valid {

			//fmt.Println(claims.UserID, claims.role)

			ctx := context.WithValue(request.Context(), userContext, "")
			next.ServeHTTP(writer, request.WithContext(ctx))
		} else {
			_, err := fmt.Fprintf(writer, " PLEASE LOGIN AGAIN")
			utils.CheckError(err)
			return
		}
	})

}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apikey := request.Header.Get("x-api-key")
		claims := &model.JWTClaims{}
		token, tokenErr := jwt.ParseWithClaims(apikey, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return mySigningKey, nil
		})

		if tokenErr != nil {
			utils.CheckError(tokenErr)
		}
		if token.Valid {
			ctx := context.WithValue(request.Context(), userContext, "")
			next.ServeHTTP(writer, request.WithContext(ctx))
		} else {
			_, err := fmt.Fprintf(writer, " PLEASE LOGIN AGAIN")
			utils.CheckError(err)
			return
		}
	})
}
