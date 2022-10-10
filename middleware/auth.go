package middleware

import (
	"GO-CRUD/config"
	"GO-CRUD/helper"
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

func VerifyToken(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		cookie, err := r.Cookie("token")
		if err != nil {
			helper.Unauthorized(w)
		} else {
			//get token value
			tokenString := cookie.Value
			claims := &config.JWTClaim{}
			//parsing token
			key := []byte(os.Getenv("JWT_KEY"))
			token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
				return key, nil
			})
			if err != nil {
				helper.Unauthorized(w)
			} else if !token.Valid {

				helper.Unauthorized(w)
			} else {

				ctx := context.WithValue(r.Context(), "student_id", claims.StudentId)
				next(w, r.WithContext(ctx), p)
			}

		}
	}
}
