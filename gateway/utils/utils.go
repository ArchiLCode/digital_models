package utils

import (
	"errors"
	"fmt"
	"gateway/internal/config"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

// Claims структура для хранения данных токена
type Claims struct {
	JTI   string `json:"jti"`
	UID   string `json:"uid"`
	Email string `json:"email"`
	AppID int    `json:"app_id"`
	jwt.StandardClaims
}

// IsValidToken проверяет валидность JWT-токена
func IsValidToken(tokenString string) bool {
	cfg := config.MustLoad()
	// Парсим токен
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи является HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Возвращаем секретный ключ в виде []byte
		return []byte(cfg.TokenKey), nil
	})

	if err != nil {
		return false
	}

	// Проверяем, является ли токен действительным
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Дополнительная проверка claims, если необходимо
		if claims.JTI == "" || claims.UID == "" || claims.Email == "" || claims.AppID <= 0 {
			return false
		}
		return true
	}

	return false
}

func GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		if err == http.ErrNoCookie {
			return "", err
		}
		return "", err
	}

	refreshToken := cookie.Value
	return refreshToken, nil
}

// ExtractToken извлекает access token из заголовка Authorization.
// Формат заголовка должен быть: "Bearer <token>".
func ExtractToken(authorizationHeader string) (string, error) {
	// Проверяем, что заголовок не пустой
	if authorizationHeader == "" {
		return "", errors.New("authorization header is empty")
	}

	// Разделяем строку по пробелам
	parts := strings.Split(authorizationHeader, " ")

	// Проверяем, что формат корректный ("Bearer <token>")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format. Expected 'Bearer <token>'")
	}

	// Возвращаем токен
	return parts[1], nil
}
