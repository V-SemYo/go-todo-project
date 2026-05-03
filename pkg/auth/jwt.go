package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("todo-jwt-key")

// HashedPassword возвращает hash пароля для встраивания в токен (JWT)
func HashedPassword(password string) string {
	return password + string(jwtSecret)
}

// GenerateToken создаёт JWT для указанного пароля
func GenerateToken(password string) (string, error) {
	claims := jwt.MapClaims{
		"passHash": HashedPassword(password),
		"exp":      time.Now().Add(8 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken проверяет токен на валидность и возвращает подтверждение, или ошибку
func ValidateToken(tokenString, currentPassword string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		storedHash, ok := claims["passHash"].(string)
		if !ok {
			return false, nil
		}
		return storedHash == HashedPassword(currentPassword), nil
	}
	return false, nil
}
