package models

import jwt "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
