package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Password string `json:"password" example:"dummy"`
	Username string `json:"username" example:"dummy"`
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

type Login struct {
	ID        int       `json:"id"`
	IdUser    int       `json:"id_user"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
}
