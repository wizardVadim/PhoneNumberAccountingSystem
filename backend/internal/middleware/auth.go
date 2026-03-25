package middleware

import (
	"context"
	"net/http"
	"phone-accounting-system/internal/models"

	"github.com/golang-jwt/jwt"
)

var JwtSecret = []byte("secret-key-phone-accounting")

type contextKey string

const UserKey contextKey = "user"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) > 7 {
			tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			user := &models.User{
				Id:    int64(claims["user_id"].(float64)),
				Login: claims["login"].(string),
				RoleId: int64(claims["role_id"].(float64)),
			}
			ctx := context.WithValue(r.Context(), UserKey, user)
			r = r.WithContext(ctx)
		}

		next(w, r)
	}
}
