package handler

import (
	"encoding/json"
	"gochicoba/helpers"
	"gochicoba/models"
	"gochicoba/service"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginHandler struct {
	loginService service.LoginService
}

func NewLoginHandler(loginService service.LoginService) LoginHandler {
	return LoginHandler{
		loginService: loginService,
	}
}

func (lh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	jwtKey := os.Getenv("SECRET_KEY")

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = lh.loginService.CheckLogin(creds)
	if err != nil {
		helpers.ErrorResponse(w, r, http.StatusUnauthorized, "login failed", err.Error())
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

	res := models.LoginResponse{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
	helpers.CustomResponse(w, r, http.StatusOK, "success", res)
}
