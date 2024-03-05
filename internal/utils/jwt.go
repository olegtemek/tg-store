package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(secret string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(300 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("cannot create access token")
	}

	return tokenString, nil
}

func GenerateRefreshToken(secret string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(60 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("cannot create refresh token")
	}

	return tokenString, nil
}

func CompareToken(tokenString string, secret string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("cannot compare token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return int(claims["userId"].(float64)), nil
	}

	return 0, fmt.Errorf("cannot parse token")
}
