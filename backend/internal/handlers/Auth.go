package handlers

import (
	"encoding/json"
	"net/http"
	"phone-accounting-system/internal/middleware"
	"phone-accounting-system/internal/models"
	"phone-accounting-system/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	UserRepo *repository.UserRepo
	Mux *http.ServeMux
}

func (h *AuthHandler) Init() {
	h.Mux.HandleFunc("POST /api/login", h.Login)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&creds)

	user := models.User{Login: creds.Login, Password: creds.Password}
	if !h.UserRepo.Auth(&user) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"login":   user.Login,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(middleware.JwtSecret)

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
