package controllers

import (
	"college-be/auth"
	"college-be/database"
	"college-be/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type TokenRequest struct {
	NIM      string `json:"nim"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meta": &models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.Db.Where("nim = ?", request.NIM).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"meta": models.Meta{false, "Invalid credentials"}})
		c.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.NIM)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "Success"},
		"data": models.LoginResponse{tokenString}})
}

func ExtractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", errors.New("unauthorized Error: Access Denied")
}

func ExtractNIM(c *gin.Context) (string, error) {
	signedToken, err := ExtractToken(c)
	if err != nil {
		return "", err
	}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&auth.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*auth.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}
	return claims.NIM, nil
}
