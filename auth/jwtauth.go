package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTClaim struct {
	NIM string `json:"nim"`
	jwt.StandardClaims
}

func GenerateJWT(nim string) (tokenString string, err error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		NIM: nim,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
