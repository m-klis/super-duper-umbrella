package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginResponse struct {
	Name    string
	Value   string
	Expires time.Time
}
