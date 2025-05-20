package jwt

import (
	"auth/internal/domain/models"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenerateRefresh создает новый refresh token в формате JWT
func GenerateRefresh(user models.User, app models.App, duration time.Duration) (string, error) {
	// Создаем новый токен с алгоритмом HS256
	token := jwt.New(jwt.SigningMethodHS256)

	// Генерируем уникальный идентификатор (jti) для токена
	jti, err := generateJTI()
	if err != nil {
		return "", err
	}

	// Устанавливаем claims для refresh token
	claims := token.Claims.(jwt.MapClaims)
	claims["jti"] = jti                             // Уникальный идентификатор токена
	claims["uid"] = user.ID                         // Идентификатор пользователя
	claims["email"] = user.Email                    // Email пользователя
	claims["app_id"] = app.ID                       // Идентификатор приложения
	claims["exp"] = time.Now().Add(duration).Unix() // Время истечения токена

	// Подписываем токен секретным ключом приложения
	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateJTI() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	jti, err := generateJTI()
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["jti"] = jti
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseAccessToken(tokenString string, app models.App) (*models.Claims, error) {
	// Создаем экземпляр парсера
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм подписи является ожидаемым
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи")
		}
		return []byte(app.Secret), nil
	})

	// Если произошла ошибка при парсинге
	if err != nil {
		return nil, err
	}

	// Проверяем, что токен валиден
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("неверный или истекший токен")
}
