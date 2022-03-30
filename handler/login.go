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

// LOGIN APP
// @Summary      For Login to API
// @Description  Login first for authorization (Bearer Authorization)
// @Description  Copy token value, then add to field authorization with format
// @Description  Example :
// @Description  Bearer ABCDEFGHIJKLMN.1234567890.XXX
// @Tags         LOGIN
// @Accept       json
// @Produce      json
// @Param note body models.Credentials true "Login"
// @Success      200  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /login [post]
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

	expirationTime := time.Now().Add(30 * time.Minute)
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
