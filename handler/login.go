package handler

import (
	"encoding/json"
	"gochicoba/helpers"
	"gochicoba/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginHandler struct {
}

func NewLoginHandler() LoginHandler {
	return LoginHandler{}
}

func (lh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	jwtKey := os.Getenv("SECRET_KEY")
	var users = map[string]string{
		"user1": "password1",
		"user2": "password2",
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })

	res := models.LoginResponse{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", res)
}
