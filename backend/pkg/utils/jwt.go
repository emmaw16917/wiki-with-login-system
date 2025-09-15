package utils

import (
	"time"
	"wiki/backend/config"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID int) (string, error) { //生成JWT
	jwtSecret := []byte(config.JWTSecret)
	expireHours := config.JWTExpireHours
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(expireHours)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (int, error) { //解析和校验JWT
	jwtSecret := []byte(config.JWTSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, err
	}
	return int(userID), nil
}
