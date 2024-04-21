package authentication

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenValidityDuration = 60 * 24 * time.Hour

func ValidateJWTToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		return "", errors.New("token has expired")
	}

	userID, username, err := GetUserByID(claims.UserID)
	if err != nil {
		return "", err
	}

	newToken, err := GenerateJWTToken(userID, username)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func GetUserByID(userId int) (int, string, error) {
	userData, ok := Users[userId]
	if ok {
		return userId, userData[1], nil
	}
	return -1, "", errors.New("user not exists")
}
