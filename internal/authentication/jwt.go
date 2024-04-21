package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("f350124f5d1497b709735c3cc842041b3534c0e1a87f0ba0abd5c37ee8369bf4")

type Claims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWTToken(userID int, username string) (string, error) {
	expirationDays := 60
	expirationTime := time.Now().Add(time.Duration(24*expirationDays) * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   fmt.Sprint(userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
